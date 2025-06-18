package core

import (
	"boilerplate/config"
	"boilerplate/pkg/infra/db"

	"github.com/sirupsen/logrus"
)

type CoreRepository struct {
}

func NewCoreRepository(conf *config.Config, dbList *db.DatabaseList, log *logrus.Logger) CoreRepository {
	return CoreRepository{
	}
}
