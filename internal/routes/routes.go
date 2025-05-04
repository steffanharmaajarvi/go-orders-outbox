package routes

import (
	"github.com/gin-gonic/gin"
	"orders.go/m/internal/controller/orders"
)

func New(
	router *gin.Engine,
) {
	orderController := orders.New()

	ordersRoutes := router.Group("/orders")
	{
		ordersRoutes.POST("/", orderController.Create)
	}
}
