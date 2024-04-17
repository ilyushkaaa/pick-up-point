package database

import (
	"context"
	"fmt"

	"homework/pkg/database/postgres/transaction_manager"
)

func New(ctx context.Context) (Database, error) {
	dsn := generateDsn()

	qp, err := transaction_manager.New(ctx, dsn)
	if err != nil {
		return nil, err
	}
	return NewDatabase(qp), nil
}

func generateDsn() string {
	connData := getConnectData()
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		connData.host, connData.port, connData.user, connData.password, connData.dbName)
}
