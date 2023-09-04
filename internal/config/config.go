package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

type Config struct {
	Postgres PostgresConfig `envconfig:"DATABASE" required:"true"`
}

func InitConfig() (Config, error) {
	_ = godotenv.Load(".env")

	var cfg Config
	if err := envconfig.Process("APP", &cfg); err != nil {
		return Config{}, errors.Wrap(err, "get config from env error")
	}

	return cfg, nil
}
