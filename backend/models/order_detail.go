package models

type OrderDetail struct {
	ID        uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	OrderID   uint    `gorm:"not null"                 json:"orderId"`
	ProductID uint    `gorm:"not null"                 json:"productId"`
	Product   Product `gorm:"foreignKey:ProductID"     json:"product"`
	Quantity  int     `gorm:"not null"                 json:"quantity"`
	Price     float64 `gorm:"not null"                 json:"price"`
}
