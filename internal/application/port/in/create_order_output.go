package in

type CreateOrderOutput struct {
	Total float64
}

func NewCreateOrderOutput(total float64) CreateOrderOutput {
	return CreateOrderOutput{
		Total: total,
	}
}
