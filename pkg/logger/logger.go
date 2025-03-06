package logger

import (
	"go.uber.org/zap"
)

var Logger *zap.Logger

func Init() {
	logger, _ := zap.NewProduction() // Use NewDevelopment() for dev
	Logger = logger
}

func Sync() {
	Logger.Sync()
}
