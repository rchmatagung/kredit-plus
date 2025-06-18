package repository

import (
	"boilerplate/config"
	"boilerplate/internal/wrapper/repository/core"
	"boilerplate/internal/wrapper/repository/general"
	"boilerplate/pkg/infra/db"

	"github.com/sirupsen/logrus"
)

type Repository struct {
	General general.GeneralRepository
	Core    core.CoreRepository
}

func NewRepository(conf *config.Config, dbList *db.DatabaseList, log *logrus.Logger) Repository {
	return Repository{
		General: general.NewGeneralRepository(conf, dbList, log),
		Core:    core.NewCoreRepository(conf, dbList, log),
	}
}
