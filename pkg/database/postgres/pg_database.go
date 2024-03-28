package database

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"homework/pkg/database/client"
)

type PGDatabase struct {
	Cluster *pgxpool.Pool
}

func NewDatabase(ctx context.Context) (*PGDatabase, error) {
	cluster, err := client.NewPool(ctx)
	if err != nil {
		return nil, err
	}
	return &PGDatabase{
		Cluster: cluster,
	}, nil
}
