package config

import (
	"flag"
	"fmt"

	"github.com/siwonpawel/cash-gopher-outside/banking/logger"
)

const (
	authAddress = "AUTH_ADDRESS"
	formatAPP   = "%s:%d"
	formatDB    = "%s:%s@tcp(%s:%d)/%s"
	appAddress  = "SERVER_ADDRESS"
	appPort     = "SERVER_PORT"
	dbUser      = "DB_USER"
	dbPassword  = "DB_PASSWORD"
	dbHost      = "DB_HOST"
	dbName      = "DB_NAME"
	dbPort      = "DB_PORT"
)

type configuration struct {
	authAddress string
	appAddress  string
	appPort     int
	dbUser      string
	dbPasswd    string
	dbHost      string
	dbPort      int
	dbName      string
}

var config configuration

func init() {
	flag.StringVar(&config.appAddress, appAddress, "localhost", "server address to listen and serve")
	flag.IntVar(&config.appPort, appPort, 8000, "server port to listen and serve")
	flag.StringVar(&config.dbUser, dbUser, "", "username for connecting to database")
	flag.StringVar(&config.dbPasswd, dbPassword, "", "password for database user")
	flag.StringVar(&config.dbHost, dbHost, "", "database host")
	flag.IntVar(&config.dbPort, dbPort, 3306, "database port")
	flag.StringVar(&config.dbName, dbName, "", "database name")
	flag.StringVar(&config.authAddress, authAddress, "localhost:8081", "auth app address")
	flag.Parse()

	if "" == config.dbName || "" == config.dbHost || "" == config.dbPasswd || "" == config.dbUser || "" == config.authAddress {
		flag.PrintDefaults()
		panic("Please provide all required variables before running application!")
	}
}

func GetAppAddress() string {
	return fmt.Sprintf(formatAPP, config.appAddress, config.appPort)
}

func GetDBAddress() string {
	logger.Info(fmt.Sprintf("Generated DB URL: "+formatDB, config.dbUser, "******", config.dbHost, config.dbPort, config.dbName))
	return fmt.Sprintf(formatDB, config.dbUser, config.dbPasswd, config.dbHost, config.dbPort, config.dbName)
}

func GetAuthAppAddress() string {
	return config.authAddress
}
