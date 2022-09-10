package service

import (
	"mapping/internal/application/port/out"
	"mapping/internal/domain"
	"time"
)

var now = time.Now

type CreateOrder interface {
	Action(*domain.Order) *domain.Order
}

type CreateOrderService struct {
	saveOrder out.SaveOrderPort
}

func NewCreateOrder(saveOrder out.SaveOrderPort) CreateOrder {
	return CreateOrderService{
		saveOrder: saveOrder,
	}
}

func (s CreateOrderService) Action(o *domain.Order) *domain.Order {
	withTax := 1.1
	o.Total = float64(o.Count*o.Price) * withTax
	o.CreatedAt = now()
	s.saveOrder.SaveOrder(o)
	return o
}
