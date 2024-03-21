package client

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

func NewPool(ctx context.Context) (*pgxpool.Pool, error) {
	dsn, err := generateDsn()
	if err != nil {
		return nil, err
	}
	pool, err := pgxpool.Connect(ctx, dsn)
	if err != nil {
		return nil, err
	}
	return pool, nil
}

func generateDsn() (string, error) {
	connData, err := getConnectData()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		connData.host, connData.port, connData.user, connData.password, connData.dbName), nil
}
