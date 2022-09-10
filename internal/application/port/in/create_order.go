package in

import "mapping/internal/domain"

type CreateOrder interface {
	Action(*domain.Order) *domain.Order
}
