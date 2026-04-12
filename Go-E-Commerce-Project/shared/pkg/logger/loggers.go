package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func Init(env string) {
	var err error
	config := zap.NewProductionConfig()

	if env == "development" {
		config = zap.NewDevelopmentConfig()
	}

	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	log, err = config.Build()
	if err != nil {
		panic(err)
	}
}

func Get() *zap.Logger {
	if log == nil {
		Init("production")
	}
	return log
}
