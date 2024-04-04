package database

import (
	"context"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type DBops interface {
	Select(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	Exec(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error)
	QueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row
	Get(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	Close() error
}

type PGDatabase struct {
	cluster *pgxpool.Pool
}

func NewDatabase(cluster *pgxpool.Pool) *PGDatabase {
	return &PGDatabase{
		cluster: cluster,
	}
}

func (db *PGDatabase) Get(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return pgxscan.Get(ctx, db.cluster, dest, query, args...)
}

func (db *PGDatabase) Select(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return pgxscan.Select(ctx, db.cluster, dest, query, args...)
}

func (db *PGDatabase) Exec(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error) {
	return db.cluster.Exec(ctx, query, args...)
}

func (db *PGDatabase) QueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row {
	return db.cluster.QueryRow(ctx, query, args...)
}
func (db *PGDatabase) Close() error {
	db.cluster.Close()
	return nil
}
