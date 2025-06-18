package middleware

import (
	"boilerplate/config"
	"boilerplate/pkg/infra/db"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

// * Init
var initData MiddlewareData

type MiddlewareData struct {
	App    *fiber.App
	DBList *db.DatabaseList
	Conf   *config.Config
	Log    *logrus.Logger
}

func InitMiddlewareConfig(app *fiber.App, dbList *db.DatabaseList, conf *config.Config, log *logrus.Logger) {
	initData = MiddlewareData{
		App:    app,
		DBList: dbList,
		Conf:   conf,
		Log:    log,
	}
}
