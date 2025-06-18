package handler

import (
	"boilerplate/config"
	"boilerplate/internal/wrapper/handler/core"
	"boilerplate/internal/wrapper/handler/general"
	"boilerplate/internal/wrapper/usecase"

	"github.com/sirupsen/logrus"
)

type Handler struct {
	General general.GeneralHandler
	Core    core.CoreHandler
}

func NewHandler(uc usecase.Usecase, conf *config.Config, log *logrus.Logger) Handler {
	return Handler{
		General: general.NewGeneralHandler(uc, conf, log),
		Core:    core.NewCoreHandler(uc, conf, log),
	}

}
