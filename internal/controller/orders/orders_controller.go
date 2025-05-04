package orders

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"orders.go/m/internal/usecases/orders"
)

type OrderController struct {
	createOrderUseCase *orders.CreateOrderUseCase
}

func New(
	createOrderUseCase *orders.CreateOrderUseCase,
) *OrderController {
	return &OrderController{
		createOrderUseCase: createOrderUseCase,
	}
}

func (controller *OrderController) Create(c *gin.Context) {
	var req orders.CreateOrderInput
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := controller.createOrderUseCase.Execute(
		c.Request.Context(),
		req,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Order created"})
	return
}
