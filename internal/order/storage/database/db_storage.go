package storage

import (
	"homework/pkg/infrastructure/database/postgres"
)

type OrderStoragePG struct {
	db database.Database
}

func New(db database.Database) *OrderStoragePG {
	return &OrderStoragePG{db: db}
}
