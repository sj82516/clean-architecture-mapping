package in

type CreateOrder interface {
	Action(command *CreateOrderCommand) *CreateOrderOutput
}
