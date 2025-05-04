package database

import (
	"database/sql"
	"orders.go/m/internal/infrastructure/config"
	"orders.go/m/internal/infrastructure/database/postgres"
)

type DatabaseProvider string

const (
	Postgres DatabaseProvider = "postgres"
)

type Database interface {
	GetConnection(connection string) (*sql.DB, error)
	CloseAll()
}

func New(provider DatabaseProvider, config config.Config) Database {
	switch provider {
	case Postgres:
		return postgres.New(config)
	default:
		return nil
	}

}
