package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID uint
	Total  float64
	Status string
	Items  []OrderItem `gorm:"foreignKey:OrderID"`
}

type OrderItem struct {
	gorm.Model
	OrderID   uint
	ProductID uint
	Quantity  int
	Price     float64
}
