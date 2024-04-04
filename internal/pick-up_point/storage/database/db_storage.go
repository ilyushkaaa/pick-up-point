package storage

import (
	"homework/pkg/database/postgres"
)

type PPStorageDB struct {
	db database.DBops
}

func New(db database.DBops) *PPStorageDB {
	return &PPStorageDB{db: db}
}
