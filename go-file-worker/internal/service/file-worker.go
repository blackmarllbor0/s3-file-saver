package service

import (
	"app/internal/models"
	"app/internal/repository"
	"app/internal/repository/postgres"
	"app/internal/repository/s3"
	"app/internal/transport/grpcClient"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
)

type fileWorkerService struct {
	fileWorkerRepo   s3.FileWorker
	fileMetadataRepo repository.FileMetadataRepo

	pool *postgres.PgxPool

	grpcClient.UnimplementedFileWorkerServer
}

func NewFileWorkerService(
	fileWorkerRepo s3.FileWorker,
	fileMetadataRepo repository.FileMetadataRepo,
	pool *postgres.PgxPool,
) grpcClient.FileWorkerServer {
	return &fileWorkerService{
		fileWorkerRepo:   fileWorkerRepo,
		fileMetadataRepo: fileMetadataRepo,
		pool:             pool,
	}
}

func (f fileWorkerService) SaveFile(ctx context.Context, file *grpcClient.File) (*grpcClient.DefaultResponse, error) {
	errResponse := &grpcClient.DefaultResponse{
		Error: &grpcClient.ResponseError{
			ErrorCode: 1,
		},
	}

	conn, err := f.pool.GetConnection()
	if err != nil {
		errResponse.Error.ErrorMsg = err.Error()
		return errResponse, err
	}

	metadata := models.FileMetadata{
		Size:        file.GetMetadata().Size,
		Name:        file.GetMetadata().Filename,
		Ext:         file.GetMetadata().Ext,
		ContentType: file.GetMetadata().ContentType,
		Path:        file.GetMetadata().Path,
		Encoding:    file.GetMetadata().Encoding,
	}

	tx, err := conn.Begin(ctx)
	if _, err := f.fileMetadataRepo.CreateFileMetadata(ctx, tx, metadata); err != nil {
		if err := tx.Rollback(ctx); err != nil {
			errResponse.Error.ErrorMsg = fmt.Sprintf(
				"service.fileWorkerService.SaveFile: error while rollback: %v",
				err,
			)
			return errResponse, fmt.Errorf(errResponse.Error.ErrorMsg)
		}

		errResponse.Error.ErrorMsg = err.Error()
		return errResponse, err
	}

	if err := f.fileWorkerRepo.SaveFile(ctx, file.GetName(), file.GetContent()); err != nil {
		if err := tx.Rollback(ctx); err != nil {
			errResponse.Error.ErrorMsg = fmt.Sprintf(
				"service.fileWorkerService.SaveFile: error while rollback: %v",
				err,
			)
			return errResponse, fmt.Errorf(errResponse.Error.ErrorMsg)
		}

		errResponse.Error.ErrorMsg = err.Error()
		return errResponse, err
	}

	if err := tx.Commit(ctx); err != nil {
		errResponse.Error.ErrorMsg = fmt.Sprintf(
			"service.fileWorkerService.SaveFile: error while commit: %v",
			err,
		)

		return errResponse, fmt.Errorf(errResponse.Error.ErrorMsg)
	}

	defer conn.Release()

	return &grpcClient.DefaultResponse{
		Msg: aws.String(fmt.Sprintf(
			"service.fileWorkerService.SaveFile: %s file successfully uploaded", file.GetName()),
		),
	}, nil
}

func (f fileWorkerService) SaveFiles(ctx context.Context, files *grpcClient.Files) (*grpcClient.DefaultResponse, error) {
	errResponse := &grpcClient.DefaultResponse{
		Error: &grpcClient.ResponseError{
			ErrorCode: 1,
		},
	}

	filesMap := make(map[string][]byte, len(files.Files))
	metadataArr := make([]models.FileMetadata, len(files.Files))
	for i, file := range files.Files {
		// there is no point in creating goroutines, because four
		// is written to the map quite quickly
		filesMap[file.GetName()] = file.GetContent()
		metadataArr[i] = models.FileMetadata{
			Size:        file.GetMetadata().Size,
			Name:        file.GetMetadata().Filename,
			Ext:         file.GetMetadata().Ext,
			ContentType: file.GetMetadata().ContentType,
			Path:        file.GetMetadata().Path,
			Encoding:    file.GetMetadata().Encoding,
		}
	}

	conn, err := f.pool.GetConnection()
	if err != nil {
		errResponse.Error.ErrorMsg = err.Error()
		return errResponse, err
	}

	tx, err := conn.Begin(ctx)
	if err := f.fileMetadataRepo.CreateManyMetadataRows(ctx, tx, metadataArr); err != nil {
		if err := tx.Rollback(ctx); err != nil {
			errResponse.Error.ErrorMsg = fmt.Sprintf(
				"service.fileWorkerService.SaveFiles: error while rollback: %v",
				err,
			)
			return errResponse, fmt.Errorf(errResponse.Error.ErrorMsg)
		}

		errResponse.Error.ErrorMsg = err.Error()
		return errResponse, err
	}

	if err := f.fileWorkerRepo.SaveFiles(ctx, filesMap); err != nil {
		errResponse.Error.ErrorMsg = err.Error()
		return errResponse, err
	}

	if err := tx.Commit(ctx); err != nil {
		errResponse.Error.ErrorMsg = fmt.Sprintf(
			"service.fileWorkerService.SaveFiles: error while commit: %v",
			err,
		)

		return errResponse, fmt.Errorf(errResponse.Error.ErrorMsg)
	}

	defer conn.Release()

	return &grpcClient.DefaultResponse{
		Msg: aws.String("service.fileWorkerService.SaveFiles: all files successfully saved"),
	}, nil
}

func (f fileWorkerService) DeleteFile(ctx context.Context, name *grpcClient.FileNames) (*grpcClient.DefaultResponse, error) {
	errResponse := &grpcClient.DefaultResponse{
		Error: &grpcClient.ResponseError{
			ErrorCode: 1,
		},
	}

	conn, err := f.pool.GetConnection()
	if err != nil {
		errResponse.Error.ErrorMsg = err.Error()
		return errResponse, err
	}

	tx, err := conn.Begin(ctx)
	if err := f.fileMetadataRepo.DeleteFiles(ctx, tx, name.GetName()); err != nil {
		if err := tx.Rollback(ctx); err != nil {
			errResponse.Error.ErrorMsg = fmt.Sprintf(
				"service.fileWorkerService.DeleteFile: error while rollback: %v",
				err,
			)
			return errResponse, fmt.Errorf(errResponse.Error.ErrorMsg)
		}

		errResponse.Error.ErrorMsg = err.Error()
		return errResponse, err
	}

	if err := f.fileWorkerRepo.DeleteFile(ctx, name.GetName()); err != nil {
		errResponse.Error.ErrorMsg = err.Error()
		return errResponse, err
	}

	if err := tx.Commit(ctx); err != nil {
		errResponse.Error.ErrorMsg = fmt.Sprintf(
			"service.fileWorkerService.DeleteFile: error while commit: %v",
			err,
		)

		return errResponse, fmt.Errorf(errResponse.Error.ErrorMsg)
	}

	defer conn.Release()

	return &grpcClient.DefaultResponse{
		Msg: aws.String(fmt.Sprintf(
			"service.fileWorkerService.DeleteFile: %v files successfully deleted",
			name.GetName(),
		)),
	}, nil
}

func (f fileWorkerService) GetFolderFiles(ctx context.Context, name *grpcClient.FolderName) (*grpcClient.Files, error) {
	//TODO implement me
	panic("implement me")
}
