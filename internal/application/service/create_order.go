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

func (s CreateOrderService) Action(command *in.CreateOrderCommand) *domain.Order {
	o := domain.Order{
		Price: command.Price,
		Count: command.Count,
	}

	withTax := 1.1
	o.Total = float64(o.Count*o.Price) * withTax

	cmd := out.NewSaveOrderCommand(o.Price, o.Count, o.Total, now())
	s.saveOrder.SaveOrder(&cmd)

	return &o
}
