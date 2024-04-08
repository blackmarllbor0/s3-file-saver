package postgres

import (
	"app/internal/models"
	"app/internal/repository"
	"context"
	"fmt"
	"strings"

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
		`INSERT INTO files.metadata (size, name, ext, content_type, path)
		 VALUES ($1, $2, $3, $4, $5)`,
		metadata.Size,
		metadata.Name,
		metadata.Ext,
		metadata.ContentType,
		metadata.Path,
	).Scan(&ID); err != nil {
		return 0, fmt.Errorf("pg: failed to insert metadata: %v", err)
	}

	return ID, nil
}

func (f fileMetadataRepo) CreateManyMetadataRows(ctx context.Context, tx pgx.Tx, metadata []models.FileMetadata) error {
	var queryBuilder strings.Builder
	queryBuilder.WriteString("INSERT INTO files.metadata (size, name, ext, content_type, path) VALUES")

	args := make([]interface{}, 0)

	for i, file := range metadata {
		num := i * 5 // num of params per file
		queryBuilder.WriteString(fmt.Sprintf("($%d, $%d, $%d, $%d, $%d),", num+1, num+2, num+3, num+4, num+5))
		args = append(args, file.Size, file.Name, file.Ext, file.ContentType, file.Path)
	}

	// delete last comma
	queryStr := queryBuilder.String()
	queryStr = queryStr[:len(queryStr)-1]

	if _, err := tx.Query(ctx, queryStr, args...); err != nil {
		return fmt.Errorf("pg: failed to batch insert metadata: %v", err)
	}

	return nil
}

func (f fileMetadataRepo) DeleteFiles(ctx context.Context, tx pgx.Tx, fileNames []string) error {
	queryBuilder := strings.Builder{}
	queryBuilder.WriteString("UPDATE files.metadata SET is_deleted = true WHERE name IN (")

	placeholders := make([]string, len(fileNames))
	args := make([]interface{}, len(fileNames))

	for i, id := range fileNames {
		placeholders[i] = fmt.Sprintf("$%d", i+1)
		args[i] = id
	}

	queryBuilder.WriteString(strings.Join(placeholders, ", "))
	queryBuilder.WriteString(")")

	if err := tx.QueryRow(ctx, queryBuilder.String(), args...); err != nil {
		return fmt.Errorf("pg: failed to batch deleted files: %v", err)
	}

	return nil
}
