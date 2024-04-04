package models

import "time"

type FileMetadata struct {
	ID               uint64    `sql:"id"`
	Size             int64     `sql:"size"`
	Name             string    `sql:"name"`
	Ext              string    `sql:"ext"`
	ContentType      string    `sql:"content_type"`
	HashSum          string    `sql:"hash_sum"`
	Path             string    `sql:"path"`
	CreatedBy        string    `sql:"created_by"`
	LastModifiedBy   string    `sql:"last_modified_by"`
	Tags             []string  `sql:"tags"`
	CreationDate     time.Time `sql:"creating_date"`
	LastModifiedDate time.Time `sql:"last_modified_time"`
}
