package config

import (
	"github.com/example/go-shopping/models"
	"github.com/example/go-shopping/seed"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("shopping.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(
		&models.Product{},
		&models.Customer{},
		&models.Order{},
		&models.OrderDetail{},
	)
	if err != nil {
		return nil, err
	}

	seed.SeedProducts(db)
	return db, nil
}
