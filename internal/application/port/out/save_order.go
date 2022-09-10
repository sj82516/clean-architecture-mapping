package out

type SaveOrderPort interface {
	SaveOrder(command *SaveOrderCommand)
}
