package postgres

import (
	"app/internal/models"
	"app/internal/repository"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type fileMetadataRepo struct {
}

func NewFileMetadataRepo() repository.FileMetadataRepo {
	return &fileMetadataRepo{}
}

func (f fileMetadataRepo) CreateFileMetadata(
	ctx context.Context,
	tx pgx.Tx,
	metadata models.FileMetadata,
) (ID uint64, err error) {
	if err := tx.QueryRow(
		ctx,
		`INSERT INTO files.files_metadata (size, name, ext, content_type, hash_sum, path, created_by)
		 VALUES ($1, $2, $3, $4, $5, $6, $7)`,
		metadata.Size,
		metadata.Name,
		metadata.Ext,
		metadata.ContentType,
		metadata.HashSum,
		metadata.Path,
		metadata.CreatedBy,
	).Scan(&ID); err != nil {
		if err := tx.Rollback(ctx); err != nil {
			return 0, fmt.Errorf("pg: failed to cancel transaction: %v", err)
		}

		return 0, fmt.Errorf("pg: failed to insert metadata: %v", err)
	}

	return ID, nil
}
