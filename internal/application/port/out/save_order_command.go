package out

import "time"

type SaveOrderCommand struct {
	Price     int
	Count     int
	Total     float64
	CreatedAt time.Time
}

func NewSaveOrderCommand(price int, count int, total float64, createdAt time.Time) SaveOrderCommand {
	return SaveOrderCommand{
		Price:     price,
		Count:     count,
		Total:     total,
		CreatedAt: createdAt,
	}
}
