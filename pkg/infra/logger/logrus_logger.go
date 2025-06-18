package logger

import (
	"boilerplate/config"
	"boilerplate/pkg/constants/general"
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
)

var logger *logrus.Logger

func LogrusGetLevel(conf *config.LogrusAccount) logrus.Level {
	switch strings.ToLower(conf.Level) {
	case "fatal":
		return logrus.FatalLevel
	case "error":
		return logrus.ErrorLevel
	case "warn":
		return logrus.WarnLevel
	case "info":
		return logrus.InfoLevel
	case "debug":
		return logrus.DebugLevel
	case "trace":
		return logrus.TraceLevel
	}
	return logrus.InfoLevel
}

func NewLogrusLogger(conf *config.LogrusAccount) *logrus.Logger {
	if logger == nil {
		logger = logrus.New()
		logger.SetFormatter(&easy.Formatter{
			TimestampFormat: general.FullTimeFormat,
			LogFormat:       fmt.Sprintf("%s\n", `[%lvl%]: "%time%" %msg%`),
		})
		logger.SetLevel(LogrusGetLevel(conf))
		return logger
	}
	return logger
}
