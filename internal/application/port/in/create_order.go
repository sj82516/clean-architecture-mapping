package in

import "mapping/internal/domain"

type CreateOrder interface {
	Action(command *CreateOrderCommand) *domain.Order
}
