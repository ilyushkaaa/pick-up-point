package database

import (
	"os"
)

const (
	OptionProd = "prod"
	OptionTest = "test"
)

type connectData struct {
	host     string
	port     string
	user     string
	password string
	dbName   string
}

func getConnectData(option string) *connectData {
	var envPrefix string
	if option == OptionTest {
		envPrefix = "TEST_"
	}
	connData := &connectData{}
	connData.password = os.Getenv(envPrefix + "DB_PASS")
	connData.user = os.Getenv(envPrefix + "DB_USER")
	connData.dbName = os.Getenv(envPrefix + "DB_NAME")
	connData.host = os.Getenv(envPrefix + "DB_HOST")
	connData.port = os.Getenv(envPrefix + "DB_PORT")
	return connData
}
