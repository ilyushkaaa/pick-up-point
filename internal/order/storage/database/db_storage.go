package storage

import database "homework/pkg/database/postgres"

type OrderStoragePG struct {
	db database.DBops
}

func New(db database.DBops) *OrderStoragePG {
	return &OrderStoragePG{db: db}
}
