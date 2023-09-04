package main

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"messenger/internal/config"
	"messenger/internal/repository"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg, err := config.InitConfig()
	if err != nil {
		panic(errors.Wrap(err, "init config error"))
	}

	connection, err := config.Connection(ctx, cfg.Postgres)
	if err != nil {
		panic(errors.Wrap(err, "init connections error"))
	}

	postgresRepository := repository.NewPostgresRepository(connection)

	fmt.Print(postgresRepository)
}
