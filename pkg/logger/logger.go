package logger


import (
	"go.uber.org/zap"
	"log"
)



func New() *zap.Logger {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("Can't load logger: %v", err)
	}

	return logger
}
