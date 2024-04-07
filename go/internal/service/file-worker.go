package service

import (
	"app/internal/repository"
	"app/internal/repository/s3"
	"app/internal/transport/grpcClient"
	"context"
)

type fileWorkerService struct {
	fileWorkerRepo   s3.FileWorker
	fileMetadataRepo repository.FileMetadataRepo

	grpcClient.UnimplementedFileWorkerServer
}

func NewFileWorkerService(
	fileWorkerRepo s3.FileWorker,
	fileMetadataRepo repository.FileMetadataRepo,
) grpcClient.FileWorkerServer {
	return &fileWorkerService{
		fileWorkerRepo:   fileWorkerRepo,
		fileMetadataRepo: fileMetadataRepo,
	}
}

func (f fileWorkerService) SaveFile(ctx context.Context, file *grpcClient.File) (*grpcClient.EmptyResponse, error) {
	return &grpcClient.EmptyResponse{}, f.fileWorkerRepo.SaveFile(ctx, file.GetName(), file.GetContent())
}

func (f fileWorkerService) SaveFiles(ctx context.Context, files *grpcClient.Files) (*grpcClient.EmptyResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (f fileWorkerService) DeleteFile(ctx context.Context, name *grpcClient.FileName) (*grpcClient.EmptyResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (f fileWorkerService) GetFolderFiles(ctx context.Context, name *grpcClient.FolderName) (*grpcClient.FileList, error) {
	//TODO implement me
	panic("implement me")
}

func (f fileWorkerService) mustEmbedUnimplementedFileWorkerServer() {
	//TODO implement me
	panic("implement me")
}

//func (fw fileWorkerService) SaveFile(ctx context.Context, file *os.File) error {
//	if err := fw.fileWorkerRepo.SaveFile(file); err != nil {
//		return fmt.Errorf("fw service: failed to save file: %v", err)
//	}
//
//	// todo: move metadata retrieval to a separate service
//
//	fileInfo, err := file.Stat()
//	if err != nil {
//		return fmt.Errorf("fw service: failed to get file metadata: %v", err)
//	}
//
//	buf := make([]byte, 512)
//	if _, err := file.Read(buf); err != nil {
//		return fmt.Errorf("fw service: failed to read file buffer: %v", err)
//	}
//
//	contentType := http.DetectContentType(buf)
//
//	hash := sha256.New()
//	if _, err := io.Copy(hash, file); err != nil {
//		return fmt.Errorf("wf service: failed to get file hash sum: %v", err)
//	}
//
//	metadata := models.FileMetadata{
//		Name:        fileInfo.Name(),
//		Size:        fileInfo.Size(),
//		Ext:         filepath.Ext(fileInfo.Name()),
//		ContentType: contentType,
//		Path:        fmt.Sprintf("/%s", fileInfo.Name()),
//		HashSum:     hex.EncodeToString(hash.Sum(nil)),
//		CreatedBy:   "User", // todo: create a creatot record
//	}
//
//	// todo: add context worker
//	if _, err := fw.fileMetadataRepo.CreateFileMetadata(ctx, nil, metadata); err != nil {
//		return fmt.Errorf("fw service: failed to insert file metadata: %v", err)
//	}
//
//	return nil
//}
