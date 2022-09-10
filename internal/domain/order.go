package domain

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	ID        int `json:"id" gorm:"autoincrement"`
	Price     int `json:"price"`
	Count     int `json:"count"`
	Total     float64
	CreatedAt time.Time
}
