package client

import (
	"os"

	"github.com/joho/godotenv"
)

const envFileName = ".env"

type connectData struct {
	host     string
	port     string
	user     string
	password string
	dbName   string
}

func getConnectData() (*connectData, error) {
	if err := godotenv.Load(envFileName); err != nil {
		return nil, err
	}
	connData := &connectData{}
	connData.password = os.Getenv("dbPass")
	connData.user = os.Getenv("dbUser")
	connData.dbName = os.Getenv("dbName")
	connData.host = os.Getenv("dbHost")
	connData.port = os.Getenv("dbPort")
	return connData, nil
}
