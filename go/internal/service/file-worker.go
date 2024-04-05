package service

import (
	"app/internal/models"
	"app/internal/repository"
	"app/internal/repository/s3"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/jackc/pgx/v5"
)

type FileWorkerService interface {
	SaveFile(ctx context.Context, tx pgx.Tx, file *os.File) error
	SaveFiles([]*os.File) error
	DeleteFile(fileName string) error
	GetFolderFiles(folderName string) ([]string, error)
}

type fileWorkerService struct {
	fileWorkerRepo   s3.FileWorker
	fileMetadataRepo repository.FileMetadataRepo
}

func NewFileWorkerService(fileWorkerRepo s3.FileWorker, fileMetadataRepo repository.FileMetadataRepo) FileWorkerService {
	return &fileWorkerService{
		fileWorkerRepo:   fileWorkerRepo,
		fileMetadataRepo: fileMetadataRepo,
	}
}

func (fw fileWorkerService) SaveFile(ctx context.Context, tx pgx.Tx, file *os.File) error {
	if err := fw.fileWorkerRepo.SaveFile(file); err != nil {
		return fmt.Errorf("fw service: failed to save file: %v", err)
	}

	// todo: move metadata retrieval to a separate service

	fileInfo, err := file.Stat()
	if err != nil {
		return fmt.Errorf("fw service: failed to get file metadata: %v", err)
	}

	buf := make([]byte, 512)
	if _, err := file.Read(buf); err != nil {
		return fmt.Errorf("fw service: failed to read file buffer: %v", err)
	}

	contentType := http.DetectContentType(buf)

	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return fmt.Errorf("wf service: failed to get file hash sum: %v", err)
	}

	metadata := models.FileMetadata{
		Name:        fileInfo.Name(),
		Size:        fileInfo.Size(),
		Ext:         filepath.Ext(fileInfo.Name()),
		ContentType: contentType,
		Path:        fmt.Sprintf("/%s", fileInfo.Name()),
		HashSum:     hex.EncodeToString(hash.Sum(nil)),
		CreatedBy:   "User", // todo: create a creatot record
	}

	// todo: add context worker
	if _, err := fw.fileMetadataRepo.CreateFileMetadata(ctx, tx, metadata); err != nil {
		return fmt.Errorf("fw service: failed to insert file metadata: %v", err)
	}

	return nil
}

func (fw fileWorkerService) SaveFiles([]*os.File) error {
	return nil
}

func (fw fileWorkerService) DeleteFile(fileName string) error {
	return nil
}

func (fw fileWorkerService) GetFolderFiles(folderName string) ([]string, error) {
	return nil, nil
}
