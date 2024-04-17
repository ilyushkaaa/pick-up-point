package storage

import (
	database "homework/pkg/database/postgres"
)

type PPStorageDB struct {
	db database.Database
}

func New(db database.Database) *PPStorageDB {
	return &PPStorageDB{
		db: db,
	}
}
