package setup

import (
	"orders.go/m/internal/infrastructure/database"
)

type App struct {
	database database.Database
}

func New() *App {
	container, err := BuildContainer()
	if err != nil {
		panic(err)
	}

	return &App{database: container.Database}
}

func (s *App) CloseDb() {
	s.database.CloseAll()
}
