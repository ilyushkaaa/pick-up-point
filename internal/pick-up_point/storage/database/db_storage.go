package storage

import "homework/pkg/database/postgres"

type PPStorageDB struct {
	db *database.PGDatabase
}

func New(db *database.PGDatabase) *PPStorageDB {
	return &PPStorageDB{db: db}
}
