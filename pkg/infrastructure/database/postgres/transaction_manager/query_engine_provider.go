package transaction_manager

import "context"

type QueryEngineProvider interface {
	GetQueryEngine(ctx context.Context) QueryEngine
	Close() error
}

func (t *TransactionManagerPGX) GetQueryEngine(ctx context.Context) QueryEngine {
	tx, ok := ctx.Value(queryKey).(QueryEngine)
	if ok {
		return tx
	}

	return t.pool
}

func (t *TransactionManagerPGX) Close() error {
	t.pool.Close()
	return nil
}
