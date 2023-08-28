package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
	"messenger/internal/infr/database"
)

type Config struct {
	Postgres database.PostgresConfig `envconfig:"MERMAN_DATABASE" required:"true" vault:"postgres/"`
}

func InitConfig() (Config, error) {
	_ = godotenv.Load(".env")

	var cfg Config
	if err := envconfig.Process("APP", &cfg); err != nil {
		return Config{}, errors.Wrap(err, "get config from env error")
	}

	return cfg, nil
}
