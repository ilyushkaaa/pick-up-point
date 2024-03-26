package storage

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"homework/pkg/database/client"
)

type PPStorageDB struct {
	cluster *pgxpool.Pool
}

func New(ctx context.Context) (*PPStorageDB, error) {
	cluster, err := client.NewPool(ctx)
	if err != nil {
		return nil, err
	}
	return &PPStorageDB{cluster: cluster}, nil
}
