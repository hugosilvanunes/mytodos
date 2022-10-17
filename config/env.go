package config

import "github.com/kelseyhightower/envconfig"

type Env struct {
	Port        string `envconfig:"PORT" default:":8080"`
	DatabaseURL string `envconfig:"DATABASE_URL"`
}

func NewEnv() (*Env, error) {
	cfg := &Env{}
	if err := envconfig.Process("", cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
