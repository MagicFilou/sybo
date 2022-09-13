package logger

import (
	cfg "sybo/configs"

	"go.uber.org/zap"
)

var logger *zap.Logger

// Logger for the service, could be prod or local
func init() {

	if cfg.GetConfig().ENV != "local" {
		logger, _ = zap.NewProduction()
	} else {
		logger, _ = zap.NewDevelopment()
	}
}

func GetLogger() *zap.Logger {
	return logger
}

func GetSugar() *zap.SugaredLogger {
	return logger.Sugar()
}
