package setup

import (
	"orders.go/m/internal/infrastructure/config"
	"orders.go/m/internal/infrastructure/database"
)

type Container struct {
	Config   config.Config
	Database database.Database
}

func BuildContainer() (*Container, error) {
	// 1. Init config
	cfg := config.NewConfig(config.ConfigProviderKoanfs)

	// 2. Init storage (можно условно: postgres или dynamo)
	databasePools := database.New(database.Postgres, cfg)

	return &Container{
		Config:   cfg,
		Database: databasePools,
	}, nil
}
