package logp

import "go.uber.org/zap"

var logger *zap.Logger

func Init() error {
	var err error
	logger, err = zap.NewProduction()
	return err
}

func L() *zap.Logger {
	return logger
}
