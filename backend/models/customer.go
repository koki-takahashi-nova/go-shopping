package models

type Customer struct {
	ID          uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string `gorm:"not null"                 json:"name"`
	Email       string `gorm:"not null"                 json:"email"`
	Address     string `                                json:"address"`
	PhoneNumber string `                                json:"phoneNumber"`
}
