package models

import "time"

type Order struct {
	ID           uint          `gorm:"primaryKey;autoIncrement" json:"id"`
	CustomerID   uint          `gorm:"not null"                 json:"customerId"`
	Customer     Customer      `gorm:"foreignKey:CustomerID"    json:"customer"`
	OrderDate    time.Time     `                                json:"orderDate"`
	OrderDetails []OrderDetail `gorm:"foreignKey:OrderID"       json:"orderDetails"`
	TotalAmount  float64       `                                json:"totalAmount"`
}

func (Order) TableName() string {
	return "orders"
}
