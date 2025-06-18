package db

import (
	"boilerplate/config"
	"context"
	"time"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoInstance struct {
	Client *mongo.Client
	DB     *mongo.Database
}

func NewMongoDBConnection(conf *config.DatabaseMongoDB, log *logrus.Logger) (*mongo.Client, *mongo.Database) {

	client, err := mongo.NewClient(options.Client().ApplyURI(conf.MongoURI))
	if err != nil {
		log.Fatal("Failed to connect MongoDB. err: " + err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal("Failed to connect MongoDB. err: " + err.Error())
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	log.Info("Connection Opened to MongoDB")

	//* Get DBName from DB source
	//dbName := utils.GetDBNameFromDriverSource(conf.MongoURI)

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
	// if conf.DriverName == "postgres" || conf.DriverName == "pgx" {
	// 	db, err = gorm.Open(postgres.Open(conf.DriverSource), &gorm.Config{
	// 		Logger: gormlog.Default.LogMode(gormlog.LogLevel(gormlog.Error)),
	// 	})
	// }

	// if err != nil {
	// 	log.Fatal("Failed to connect database " + dbName + ", err: " + err.Error())
	// }

	// log.Info("Connection Opened to Database " + dbName)
	// MI = MongoInstance{
	// 	Client: client,
	// 	DB:     client.Database(conf.DB),
	// }
	return client, client.Database(conf.DB)
}
