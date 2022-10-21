package config

import "github.com/kelseyhightower/envconfig"

type Env struct {
	Environment  string `envconfig:"ENV" default:"dev"`
	Port         string `envconfig:"PORT" default:":8080"`
	DatabaseURL  string `envconfig:"DSN"`
	JWTSecretKey string `envconfig:"JWT_SECRET_KEY"`
}

func NewEnv() (*Env, error) {
	cfg := &Env{}
	if err := envconfig.Process("", cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
