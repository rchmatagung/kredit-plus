package usecase

import (
	"boilerplate/config"
	"boilerplate/internal/wrapper/repository"
	"boilerplate/internal/wrapper/usecase/core"
	"boilerplate/internal/wrapper/usecase/general"
	"boilerplate/pkg/infra/db"

	"github.com/sirupsen/logrus"
)

type Usecase struct {
	General general.GeneralUsecase
	Core    core.CoreUsecase
}

func NewUsecase(repo repository.Repository, conf *config.Config, dbList *db.DatabaseList, log *logrus.Logger) Usecase {
	return Usecase{
		General: general.NewGeneralUsecase(repo, conf, dbList, log),
		Core:    core.NewCoreUsecase(repo, conf, dbList, log),
	}
}
