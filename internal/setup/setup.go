package setup

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"orders.go/m/internal/infrastructure/database"
	"orders.go/m/internal/repository/orders"
	"orders.go/m/internal/repository/outbox"
	"orders.go/m/internal/routes"
	uow2 "orders.go/m/internal/uow"
	orders2 "orders.go/m/internal/usecases/orders"
)

type App struct {
	database database.Database
	router   *gin.Engine
}

func New() *App {
	container, err := BuildContainer()
	if err != nil {
		panic(err)
	}

	orderDb, err := container.Database.GetConnection("orders")
	if err != nil {
		panic(err)
	}

	router := container.Engine

	uow := uow2.NewUnitOfWork(orderDb)

	ordersRepo := orders.NewOrdersRepo(orderDb)
	outboxRepo := outbox.NewOutboxRepo(orderDb)

	createOrderUseCase := orders2.NewCreateOrderUseCase(ordersRepo, uow, outboxRepo)

	routes.NewOrdersRoutes(router, createOrderUseCase)

	for _, route := range router.Routes() {
		fmt.Printf("%s %s → %s\n", route.Method, route.Path, route.Handler)
	}

	return &App{
		database: container.Database,
		router:   router,
	}
}

func (s *App) CloseDb() {
	s.database.CloseAll()
}

func (s *App) Run() error {
	return s.router.Run()
}
