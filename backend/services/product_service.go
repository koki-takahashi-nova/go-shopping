package services

import (
	"strings"

	"github.com/example/go-shopping/models"
	"gorm.io/gorm"
)

const (
	LowPriceMax  = 10000.0
	MidPriceMin  = 10001.0
	MidPriceMax  = 50000.0
	HighPriceMin = 50001.0
)

type ProductService struct {
	db *gorm.DB
}

func NewProductService(db *gorm.DB) *ProductService {
	return &ProductService{db: db}
}

func (s *ProductService) GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	result := s.db.Find(&products)
	return products, result.Error
}

func (s *ProductService) GetProductByID(id uint) (*models.Product, error) {
	var product models.Product
	result := s.db.First(&product, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &product, nil
}

func (s *ProductService) GetProductsByCategory(category string) ([]models.Product, error) {
	var products []models.Product
	result := s.db.Where("category = ?", category).Find(&products)
	return products, result.Error
}

func (s *ProductService) SearchProducts(keyword *string, minPrice, maxPrice *float64) ([]models.Product, error) {
	var products []models.Product

	hasKeyword := keyword != nil && strings.TrimSpace(*keyword) != ""
	hasMin := minPrice != nil
	hasMax := maxPrice != nil

	q := s.db.Model(&models.Product{})

	if hasKeyword {
		kw := "%" + strings.TrimSpace(*keyword) + "%"
		q = q.Where("name LIKE ?", kw)
	}
	if hasMin {
		q = q.Where("price >= ?", *minPrice)
	}
	if hasMax {
		q = q.Where("price <= ?", *maxPrice)
	}

	result := q.Find(&products)
	return products, result.Error
}
