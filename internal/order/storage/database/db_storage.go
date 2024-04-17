package storage

import (
	"homework/pkg/database/postgres"
)

type OrderStoragePG struct {
	db database.Database
}

func New(db database.Database) *OrderStoragePG {
	return &OrderStoragePG{db: db}
}
