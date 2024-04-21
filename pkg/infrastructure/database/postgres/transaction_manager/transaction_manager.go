package transaction_manager

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/multierr"
)

type key string

const queryKey key = "transaction"

type TransactionManager interface {
	RunSerializable(ctx context.Context, f func(ctxTX context.Context) error) error
	RunReadCommitted(ctx context.Context, f func(ctxTX context.Context) error) error
	RunRepeatableRead(ctx context.Context, f func(ctxTX context.Context) error) error
	RunReadUnCommitted(ctx context.Context, f func(ctxTX context.Context) error) error
}

type TransactionManagerPGX struct {
	pool *pgxpool.Pool
}

func NewTM(ctx context.Context, dsn string) (*TransactionManagerPGX, error) {
	pool, err := pgxpool.Connect(ctx, dsn)
	if err != nil {
		return nil, err
	}
	return &TransactionManagerPGX{
		pool: pool,
	}, nil
}

func (t *TransactionManagerPGX) RunTransaction(ctx context.Context, isoLevel pgx.TxIsoLevel, f func(ctxTX context.Context) error) error {
	tx, err := t.pool.BeginTx(ctx, pgx.TxOptions{
		IsoLevel:   isoLevel,
		AccessMode: pgx.ReadWrite,
	})
	if err != nil {
		return err
	}

	if err = f(context.WithValue(ctx, queryKey, tx)); err != nil {
		errRollback := tx.Rollback(ctx)
		return multierr.Combine(err, errRollback)
	}

	if err = tx.Commit(ctx); err != nil {
		return multierr.Combine(err, tx.Rollback(ctx))
	}

	return nil
}

func (t *TransactionManagerPGX) RunSerializable(ctx context.Context, f func(ctxTX context.Context) error) error {
	return t.RunTransaction(ctx, pgx.RepeatableRead, f)
}

func (t *TransactionManagerPGX) RunReadCommitted(ctx context.Context, f func(ctxTX context.Context) error) error {
	return t.RunTransaction(ctx, pgx.Serializable, f)
}

func (t *TransactionManagerPGX) RunRepeatableRead(ctx context.Context, f func(ctxTX context.Context) error) error {
	return t.RunTransaction(ctx, pgx.ReadCommitted, f)
}

func (t *TransactionManagerPGX) RunReadUnCommitted(ctx context.Context, f func(ctxTX context.Context) error) error {
	return t.RunTransaction(ctx, pgx.ReadCommitted, f)
}
