package storage

import database "homework/pkg/database/postgres"

type OrderStoragePG struct {
	db *database.PGDatabase
}

func New(db *database.PGDatabase) *OrderStoragePG {
	return &OrderStoragePG{db: db}
}
