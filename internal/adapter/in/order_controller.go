package in

import (
	"mapping/internal/adapter/out"
	"mapping/internal/application/service"
	"mapping/internal/domain"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type OrderController struct {
	repo out.OrderRepository
}

func NewOrderController(db *gorm.DB, router gin.IRouter) OrderController {
	var o = OrderController{
		repo: out.NewOrderRepository(db),
	}

	router.POST("/orders", o.CreateOrder)
	return o
}

func (c *OrderController) CreateOrder(ctx *gin.Context) {
	var srv = service.NewCreateOrder(c.repo)

	type CreateOrderRequest struct {
		Price int `json:"price"`
		Count int `json:"count"`
	}

	var o CreateOrderRequest
	if err := ctx.ShouldBindJSON(&o); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	order := domain.Order{Price: o.Price, Count: o.Count}
	srv.Action(&order)

	ctx.JSON(http.StatusOK, gin.H{"total": order.Total})
}
