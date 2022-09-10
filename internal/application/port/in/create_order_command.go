package in

import "errors"

type CreateOrderCommand struct {
	Price int
	Count int
}

func NewCreateOrderCommand(price int, count int) (CreateOrderCommand, error) {
	if price < 0 || count < 0 {
		return CreateOrderCommand{}, errors.New("params error")
	}

	return CreateOrderCommand{
		Price: price,
		Count: count,
	}, nil
}
