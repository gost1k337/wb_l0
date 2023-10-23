package logging

import (
	"github.com/gost1k337/wb_demo/subscriber/config"
	"go.uber.org/zap"
)

type Logger struct {
	*zap.SugaredLogger
}

func NewLogger(cfg *config.Config) Logger {
	zapConfig := zap.NewDevelopmentConfig()

	logger, _ := zapConfig.Build()
	sugar := logger.Sugar()

	return Logger{sugar}
}
