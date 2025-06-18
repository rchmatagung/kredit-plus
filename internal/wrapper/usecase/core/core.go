package core

import (
	"boilerplate/config"
	"boilerplate/internal/wrapper/repository"
	"boilerplate/pkg/infra/db"

	"github.com/sirupsen/logrus"
)

type CoreUsecase struct {
}

func NewCoreUsecase(repo repository.Repository, conf *config.Config, dbList *db.DatabaseList, log *logrus.Logger) CoreUsecase {
	return CoreUsecase{
	}
}
