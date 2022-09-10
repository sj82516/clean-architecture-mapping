package out

import (
	"fmt"
	"mapping/internal/domain"
	"time"

	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	db.AutoMigrate(&domain.Order{})
	return OrderRepository{
		db: db,
	}
}

type OrderDAO struct {
	Id        int `gorm:"autoincrement"`
	Price     int
	Count     int
	Total     float64
	CreatedAt time.Time
}

func (r OrderRepository) SaveOrder(o *domain.Order) {
	orderDao := OrderDAO{
		Price: o.Price,
		Count: o.Count,
		Total: o.Total,
	}

	result := r.db.Create(&orderDao)
	fmt.Println(result.Error)
}
