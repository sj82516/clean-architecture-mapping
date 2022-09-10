package out

import (
	"mapping/internal/domain"
)

type SaveOrderPort interface {
	SaveOrder(o *domain.Order)
}
