package domain

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	ID    int
	Price int
	Count int
	Total float64
}
