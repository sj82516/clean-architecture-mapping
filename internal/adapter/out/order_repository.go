package out

import (
	"fmt"
	"mapping/internal/domain"

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

func (r OrderRepository) SaveOrder(o *domain.Order) {
	result := r.db.Create(&o)
	fmt.Println(result.Error)
}
