package db

import (
	"boilerplate/config"
	"boilerplate/pkg/utils"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormlog "gorm.io/gorm/logger"
)

type DatabaseList struct {
	DatabaseApp *gorm.DB
}

func NewGORMConnection(conf *config.DatabaseAccount, log *logrus.Logger) *gorm.DB {

	var db *gorm.DB
	var err error

	//* Get DBName from DB source
	dbName := utils.GetDBNameFromDriverSource(conf.DriverSource)

	// //* GORM Configuration
	// gormConf := &gorm.Config{
	// 	//* Disable gorm log
	// 	Logger: gormlog.Default.LogMode(gormlog.LogLevel(gormlog.Error)),
	// 	// Logger: gormlog.Default.LogMode(gormlog.LogLevel(gormlog.Info)),
	// 	// Logger: gormlog.Default.LogMode(gormlog.Silent),
	// 	//* Table name is singular
	// 	NamingStrategy: schema.NamingStrategy{
	// 		SingularTable: true,
	// 	},
	// 	//* Skip default gorm tx
	// 	// SkipDefaultTransaction: true,
	// }

	//* Open Connection depend on driver
	if conf.DriverName == "postgres" || conf.DriverName == "pgx" {
		db, err = gorm.Open(postgres.Open(conf.DriverSource), &gorm.Config{
			Logger: gormlog.Default.LogMode(gormlog.LogLevel(gormlog.Error)),
		})
	}

	if err != nil {
		log.Fatal("Failed to connect database " + dbName + ", err: " + err.Error())
	}

	log.Info("Connection Opened to Database " + dbName)
	return db
}
