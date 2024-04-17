package transaction_manager

import (
	"context"
	"fmt"
)

func New(ctx context.Context) (*TransactionManagerPGX, error) {
	dsn := generateDsn()
	return NewTM(ctx, dsn)

}

func generateDsn() string {
	connData := getConnectData()
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		connData.host, connData.port, connData.user, connData.password, connData.dbName)
}
