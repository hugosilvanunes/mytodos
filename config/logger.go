package config

import (
	"go.uber.org/zap"
)

func NewLogger(env string) (*zap.Logger, error) {
	logger, err := zap.NewDevelopment()
	if env == "production" {
		logger, err = zap.NewProduction()
	}

	if err != nil {
		return nil, err
	}

	return logger, nil
}
