package database

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
	connData := &connectData{}
	connData.password = os.Getenv("DB_PASS")
	connData.user = os.Getenv("DB_USER")
	connData.dbName = os.Getenv("DB_NAME")
	connData.host = os.Getenv("DB_HOST")
	connData.port = os.Getenv("DB_PORT")
	return connData
}
