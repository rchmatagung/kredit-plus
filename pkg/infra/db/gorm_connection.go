package db

import (
	"boilerplate/config"
	"boilerplate/pkg/utils"
	"database/sql"

	_ "github.com/lib/pq"

	"github.com/sirupsen/logrus"
)

type DatabaseList struct {
	DatabaseApp *sql.DB
}

func NewGORMConnection(conf *config.DatabaseAccount, log *logrus.Logger) *sql.DB {
	var db *sql.DB
	var err error

	dbName := utils.GetDBNameFromDriverSource(conf.DriverSource)

	if conf.DriverName == "postgres" || conf.DriverName == "pgx" {
		db, err = sql.Open(conf.DriverName, conf.DriverSource)
		if err != nil {
			log.Fatal("Failed to connect database " + dbName + ", err: " + err.Error())
		}

		if err = db.Ping(); err != nil {
			log.Fatal("Failed to connect database " + dbName + ", err: " + err.Error())
		}
	}

	log.Info("Connection Opened to Database " + dbName)
	return db
}
