package main

import (
	"boilerplate/config"
	"boilerplate/internal/server"
	"boilerplate/pkg/infra/db"
	"boilerplate/pkg/infra/logger"
)

func main() {

	//* ====================== Config ======================

	conf := config.InitConfig("local")

	//* ====================== Logger ======================

	//* Loggrus
	appLogger := logger.NewLogrusLogger(&conf.Logger.Logrus)

	//* Grafana Loki
	if conf.Grafana.IsActive {
		if conf.App.Env != "local" {
			err := logger.InitLoki(conf, appLogger)
			if err != nil {
				appLogger.Errorf("Grafana Loki err: %s", err.Error())
			}
		}
	}

	//* ====================== Connection DB ======================

	//var dbList db.MongoInstance

	var dbList db.DatabaseList
	dbList.DatabaseApp = db.NewGORMConnection(&conf.Connection.DatabaseApp, appLogger)
	//? Wab Fondasi Mongo DB

	//* ====================== Running Server ======================

	server.Run(conf, &dbList, appLogger)
}
