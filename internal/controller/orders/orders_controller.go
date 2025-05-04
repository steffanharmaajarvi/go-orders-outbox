package orders

import (
	"github.com/gin-gonic/gin"
)

type OrderController struct{}

func New() *OrderController {
	return &OrderController{}
}

func (controller *OrderController) Create(ctx *gin.Context) {

}
