package general

import (
	"boilerplate/config"
	"boilerplate/internal/wrapper/usecase"
	notfound "boilerplate/internal/general/notfound/delivery"
	root "boilerplate/internal/general/root/delivery"

	"github.com/sirupsen/logrus"
)

type GeneralHandler struct {
	Root     root.RootHandler
	NotFound notfound.NotFoundHandler
}

func NewGeneralHandler(uc usecase.Usecase, conf *config.Config, log *logrus.Logger) GeneralHandler {
	return GeneralHandler{
		Root:     root.NewRootHandler(uc, conf, log),
		NotFound: notfound.NewNotFoundHandler(uc, conf, log),
	}
}
