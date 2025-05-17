package setup

import (
	"github.com/gin-gonic/gin"
	"orders.go/m/internal/infrastructure/config"
	"orders.go/m/internal/infrastructure/database"
)

type Container struct {
	Config   config.Config
	Database database.Database
	Engine   *gin.Engine
}

func BuildContainer() (*Container, error) {
	// 1. Init config
	cfg := config.NewConfig(config.ConfigProviderKoanfs)

	// 2. Init storage (можно условно: postgres или dynamo)
	databasePools := database.New(database.Postgres, cfg)

	engine := gin.Default()

	return &Container{
		Config:   cfg,
		Database: databasePools,
		Engine:   engine,
	}, nil
}

func BuildConsumerContainer() (*Container, error) {
	// 1. Init config
	cfg := config.NewConfig(config.ConfigProviderKoanfs)

	// 2. Init storage (можно условно: postgres или dynamo)
	databasePools := database.New(database.Postgres, cfg)

	return &Container{
		Config:   cfg,
		Database: databasePools,
	}, nil
}
