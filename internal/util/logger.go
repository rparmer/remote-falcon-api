package util

import (
	"log"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.SugaredLogger

func NewLogger() *zap.SugaredLogger {
	loggerConfig := zap.NewProductionConfig()
	loggerConfig.EncoderConfig.TimeKey = "timestamp"
	loggerConfig.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)

	zlogger, err := loggerConfig.Build()
	if err != nil {
		log.Fatal(err)
	}

	logger = zlogger.Sugar()
	return logger

}

func GetLogger() *zap.SugaredLogger {
	return logger
}
