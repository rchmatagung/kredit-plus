package general

import (
	"boilerplate/config"
	"boilerplate/pkg/infra/db"

	"github.com/sirupsen/logrus"
)

type GeneralRepository struct {
}

func NewGeneralRepository(conf *config.Config, dbList *db.DatabaseList, log *logrus.Logger) GeneralRepository {
	return GeneralRepository{}
}
