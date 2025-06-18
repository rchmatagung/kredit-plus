package core

import (
	"boilerplate/config"
	"boilerplate/internal/wrapper/usecase"

	"github.com/sirupsen/logrus"
)

type CoreHandler struct {
}

func NewCoreHandler(uc usecase.Usecase, conf *config.Config, log *logrus.Logger) CoreHandler {
	return CoreHandler{
	}
}
