package models

import "time"

type FileMetadata struct {
	ID               uint64    `sql:"id"`
	Size             int64     `sql:"size"`
	Name             string    `sql:"name"`
	Ext              string    `sql:"ext"`
	ContentType      string    `sql:"content_type"`
	Path             string    `sql:"path"`
	Encoding         string    `sql:"encoding"`
	IsDeleted        bool      `sql:"is_deleted"`
	CreationDate     time.Time `sql:"creating_date"`
	LastModifiedDate time.Time `sql:"last_modified_time"`
}
