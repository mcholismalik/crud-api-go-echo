package database

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	dbConnections map[string]*gorm.DB
)

func Init() {
	dbConfigurations := map[string]Db{
		"SAMPLE1": &dbPostgreSQL{
			db: db{
				Host: os.Getenv("DB_HOST_SAMPLE1"),
				User: os.Getenv("DB_USER_SAMPLE1"),
				Pass: os.Getenv("DB_PASS_SAMPLE1"),
				Port: os.Getenv("DB_PORT_SAMPLE1"),
				Name: os.Getenv("DB_NAME_SAMPLE1"),
			},
			SslMode: os.Getenv("DB_SSLMODE_SAMPLE1"),
			Tz:      os.Getenv("DB_TZ_SAMPLE1"),
		},
	}

	dbConnections = make(map[string]*gorm.DB)
	for k, v := range dbConfigurations {
		db, err := v.Init()
		if err != nil {
			panic(fmt.Sprintf("Failed to connect to database %s", k))
		}
		dbConnections[k] = db
		logrus.Info(fmt.Sprintf("Successfully connected to database %s", k))
	}
}

func Connection(name string) (*gorm.DB, error) {
	if dbConnections[strings.ToUpper(name)] == nil {
		return nil, errors.New("Connection is undefined")
	}
	return dbConnections[strings.ToUpper(name)], nil
}
