package routes

import (
	"github.com/gin-gonic/gin"
	"orders.go/m/internal/controller/orders"
	orders2 "orders.go/m/internal/usecases/orders"
)

func NewOrdersRoutes(
	router *gin.Engine,
	createOrderUseCase *orders2.CreateOrderUseCase,
) {
	orderController := orders.New(createOrderUseCase)

	ordersRoutes := router.Group("/orders")
	{
		ordersRoutes.POST("/", orderController.Create)
	}
}
