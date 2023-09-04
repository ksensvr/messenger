package config

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
)

type PostgresConfig struct {
	Server       string `envconfig:"SERVER"`
	Port         uint16 `envconfig:"PORT"`
	User         string `envconfig:"USER"`
	Password     string `envconfig:"PASSWORD"`
	DatabaseName string `envconfig:"DB_NAME"`
}

func Connection(ctx context.Context, cfg PostgresConfig) (*pgxpool.Pool, error) {
	cfgStr := getPostgresDsn(cfg)

	cfgParsed, err := pgxpool.ParseConfig(cfgStr)
	if err != nil {
		return nil, err
	}

	pgConn, err := pgxpool.ConnectConfig(ctx, cfgParsed)
	if err != nil {
		return nil, err
	}

	return pgConn, nil
}

func getPostgresDsn(cfg PostgresConfig) string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		cfg.Server,
		cfg.User,
		cfg.Password,
		cfg.DatabaseName,
		cfg.Port,
	)
}
