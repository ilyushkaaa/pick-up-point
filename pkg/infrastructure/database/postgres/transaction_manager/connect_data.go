package transaction_manager

import (
	"os"
)

type connectData struct {
	host     string
	port     string
	user     string
	password string
	dbName   string
}

func getConnectData() *connectData {
	return &connectData{
		host:     os.Getenv("DB_HOST"),
		port:     os.Getenv("DB_PORT"),
		user:     os.Getenv("DB_USER"),
		password: os.Getenv("DB_PASS"),
		dbName:   os.Getenv("DB_NAME"),
	}
}
