package main

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"messenger/internal/config"
	"messenger/internal/repository"
)

func main() {
	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg, err := config.InitConfig()
	if err != nil {
		panic(errors.Wrap(err, "init config error"))
	}

	postgresRepository, err := repository.NewPostgresRepository(cfg.Postgres)
	if err != nil {
		panic(errors.Wrap(err, "init postgres repo error"))
	}

	fmt.Print(postgresRepository)
}
