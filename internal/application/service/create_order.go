package service

import (
	"mapping/internal/application/port/in"
	"mapping/internal/application/port/out"
	"mapping/internal/domain"
	"time"
)

var now = time.Now

type CreateOrderService struct {
	saveOrder out.SaveOrderPort
}

func NewCreateOrder(saveOrder out.SaveOrderPort) in.CreateOrder {
	return CreateOrderService{
		saveOrder: saveOrder,
	}
}

func (s CreateOrderService) Action(o *domain.Order) *domain.Order {
	withTax := 1.1
	o.Total = float64(o.Count*o.Price) * withTax
	s.saveOrder.SaveOrder(o)
	return o
}
