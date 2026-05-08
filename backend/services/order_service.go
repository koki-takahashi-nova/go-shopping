package services

import (
	"time"

	"github.com/example/go-shopping/models"
	"gorm.io/gorm"
)

type CartItem struct {
	ProductID uint
	Quantity  int
}

type OrderService struct {
	db *gorm.DB
}

func NewOrderService(db *gorm.DB) *OrderService {
	return &OrderService{db: db}
}

func (s *OrderService) CreateOrder(customer models.Customer, cartItems []CartItem) (*models.Order, error) {
	var order *models.Order

	err := s.db.Transaction(func(tx *gorm.DB) error {
		if result := tx.Create(&customer); result.Error != nil {
			return result.Error
		}

		order = &models.Order{
			CustomerID:  customer.ID,
			OrderDate:   time.Now(),
			OrderDetails: []models.OrderDetail{},
		}

		var totalAmount float64
		for _, item := range cartItems {
			var product models.Product
			if result := tx.First(&product, item.ProductID); result.Error != nil {
				return result.Error
			}
			subtotal := product.Price * float64(item.Quantity)
			totalAmount += subtotal
			order.OrderDetails = append(order.OrderDetails, models.OrderDetail{
				ProductID: product.ID,
				Quantity:  item.Quantity,
				Price:     product.Price,
			})
		}
		order.TotalAmount = totalAmount

		if result := tx.Create(order); result.Error != nil {
			return result.Error
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return order, nil
}
