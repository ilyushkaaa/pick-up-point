package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

func NewDB(ctx context.Context, option string) (*PGDatabase, error) {
	dsn := generateDsn(option)
	pool, err := pgxpool.Connect(ctx, dsn)
	if err != nil {
		return nil, err
	}
	return NewDatabase(pool), nil
}

func generateDsn(option string) string {
	connData := getConnectData(option)
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		connData.host, connData.port, connData.user, connData.password, connData.dbName)
}
