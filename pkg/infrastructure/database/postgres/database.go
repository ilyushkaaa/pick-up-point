package database

import (
	"context"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"homework/pkg/infrastructure/database/postgres/transaction_manager"
)

type Database interface {
	Select(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	Exec(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error)
	QueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row
	Get(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	Close() error
}

type PgxDatabase struct {
	transaction_manager.QueryEngineProvider
}

func New(qp transaction_manager.QueryEngineProvider) *PgxDatabase {
	return &PgxDatabase{
		QueryEngineProvider: qp,
	}
}

func (db *PgxDatabase) Get(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	engine := db.QueryEngineProvider.GetQueryEngine(ctx)
	return pgxscan.Get(ctx, engine, dest, query, args...)
}

func (db *PgxDatabase) Select(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	engine := db.QueryEngineProvider.GetQueryEngine(ctx)
	return pgxscan.Select(ctx, engine, dest, query, args...)
}

func (db *PgxDatabase) Exec(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error) {
	engine := db.QueryEngineProvider.GetQueryEngine(ctx)
	return engine.Exec(ctx, query, args...)
}

func (db *PgxDatabase) QueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row {
	engine := db.QueryEngineProvider.GetQueryEngine(ctx)
	return engine.QueryRow(ctx, query, args...)
}

func (db *PgxDatabase) Close() error {
	return db.QueryEngineProvider.Close()
}
