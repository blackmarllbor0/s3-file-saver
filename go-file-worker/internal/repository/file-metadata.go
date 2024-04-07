package repository

import (
	"app/internal/models"
	"context"
	"github.com/jackc/pgx/v5"
)

type FileMetadataRepo interface {
	CreateFileMetadata(ctx context.Context, tx pgx.Tx, metadata models.FileMetadata) (ID uint64, err error)
}
