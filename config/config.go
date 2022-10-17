package config

import "github.com/jmoiron/sqlx"

type Config struct {
	ENV Env
	DB  *sqlx.DB
}

func New() (*Config, error) {
	env, err := NewEnv()
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
	}, nil
}
