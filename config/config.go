package config

import (
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type Config struct {
	ENV Env
	DB  *sqlx.DB
	Log *zap.Logger
}

func New() (*Config, error) {
	env, err := NewEnv()
	if err != nil {
		return nil, err
	}

	logger, err := NewLogger(env.Environment)
	if err != nil {
		return nil, err
	}

	db, err := NewDB(env.DatabaseURL)
	if err != nil {
		return nil, err
	}

	return &Config{
		ENV: *env,
		DB:  db,
		Log: logger,
	}, nil
}
