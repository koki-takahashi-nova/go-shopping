package handlers

import (
	"net/http"

	"github.com/example/go-shopping/models"
	"github.com/example/go-shopping/services"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

type OrderHandler struct {
	orderSvc    *services.OrderService
	cartHandler *CartHandler
}

func NewOrderHandler(orderSvc *services.OrderService, cartHandler *CartHandler) *OrderHandler {
	return &OrderHandler{orderSvc: orderSvc, cartHandler: cartHandler}
}

type PlaceOrderRequest struct {
	Name        string `json:"name"        binding:"required"`
	Email       string `json:"email"       binding:"required,email"`
	Address     string `json:"address"     binding:"required"`
	PhoneNumber string `json:"phoneNumber" binding:"required"`
}

func (h *OrderHandler) PlaceOrder(c *gin.Context) {
	var req PlaceOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	session, data, err := h.cartHandler.GetSession(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(data) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cart is empty"})
		return
	}

	customer := models.Customer{
		Name:        req.Name,
		Email:       req.Email,
		Address:     req.Address,
		PhoneNumber: req.PhoneNumber,
	}

	cartItems := make([]services.CartItem, 0, len(data))
	for productID, qty := range data {
		cartItems = append(cartItems, services.CartItem{
			ProductID: productID,
			Quantity:  qty,
		})
	}

	order, err := h.orderSvc.CreateOrder(customer, cartItems)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := h.cartHandler.ClearCart(c, session); err != nil {
		// カートのクリアに失敗しても注文は成功しているので警告のみ
		c.JSON(http.StatusOK, gin.H{"order": order, "warning": "cart clear failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"order": order})
}

// GetSession の戻り値順を合わせるためのラッパー (session, data, err)
func (h *CartHandler) getSessionForOrder(c *gin.Context) (*sessions.Session, cartData, error) {
	data, session, err := h.getCart(c)
	return session, data, err
}
