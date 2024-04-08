package repository

import (
	"app/internal/models"
	"context"
	"github.com/jackc/pgx/v5"
)

type FileMetadataRepo interface {
	CreateFileMetadata(ctx context.Context, tx pgx.Tx, metadata models.FileMetadata) (ID uint64, err error)
	CreateManyMetadataRows(ctx context.Context, tx pgx.Tx, metadata []models.FileMetadata) error
	DeleteFiles(ctx context.Context, tx pgx.Tx, fileNames []string) error
}
