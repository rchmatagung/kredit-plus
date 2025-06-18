package general

import (
	"boilerplate/config"
	"boilerplate/internal/wrapper/repository"
	"boilerplate/pkg/infra/db"

	"github.com/sirupsen/logrus"
)

type GeneralUsecase struct {
}

func NewGeneralUsecase(repo repository.Repository, conf *config.Config, dbList *db.DatabaseList, log *logrus.Logger) GeneralUsecase {
	return GeneralUsecase{}
}
