package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/example/go-shopping/models"
	"github.com/example/go-shopping/services"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"gorm.io/gorm"
)

const cartSessionKey = "cart"

// cartData はセッションに保存するカートデータ (productID -> quantity)
type cartData map[uint]int

type CartHandler struct {
	store *sessions.CookieStore
	db    *gorm.DB
}

func NewCartHandler(store *sessions.CookieStore, db *gorm.DB) *CartHandler {
	return &CartHandler{store: store, db: db}
}

type CartItemResponse struct {
	Product  models.Product `json:"product"`
	Quantity int            `json:"quantity"`
	Subtotal float64        `json:"subtotal"`
}

type CartResponse struct {
	Items []CartItemResponse `json:"items"`
	Total float64            `json:"total"`
}

func (h *CartHandler) getCart(c *gin.Context) (cartData, *sessions.Session, error) {
	session, err := h.store.Get(c.Request, cartSessionKey)
	if err != nil {
		return cartData{}, nil, err
	}

	data := cartData{}
	if raw, ok := session.Values["items"]; ok {
		if b, ok := raw.([]byte); ok {
			_ = json.Unmarshal(b, &data)
		}
	}
	return data, session, nil
}

func (h *CartHandler) saveCart(c *gin.Context, session *sessions.Session, data cartData) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}
	session.Values["items"] = b
	return session.Save(c.Request, c.Writer)
}

func (h *CartHandler) buildResponse(data cartData) (CartResponse, error) {
	var productSvc services.ProductService
	_ = productSvc // not used directly; query DB directly here

	resp := CartResponse{Items: []CartItemResponse{}}
	var total float64

	for productID, qty := range data {
		var product models.Product
		if result := h.db.First(&product, productID); result.Error != nil {
			continue
		}
		subtotal := product.Price * float64(qty)
		total += subtotal
		resp.Items = append(resp.Items, CartItemResponse{
			Product:  product,
			Quantity: qty,
			Subtotal: subtotal,
		})
	}
	resp.Total = total
	return resp, nil
}

func (h *CartHandler) GetCart(c *gin.Context) {
	data, _, err := h.getCart(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	resp, err := h.buildResponse(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *CartHandler) AddToCart(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product id"})
		return
	}
	productID := uint(id)

	var product models.Product
	if result := h.db.First(&product, productID); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
		return
	}

	data, session, err := h.getCart(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	data[productID]++

	if err := h.saveCart(c, session, data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.buildResponse(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *CartHandler) RemoveFromCart(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product id"})
		return
	}
	productID := uint(id)

	data, session, err := h.getCart(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	delete(data, productID)

	if err := h.saveCart(c, session, data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.buildResponse(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *CartHandler) ClearCart(c *gin.Context, session *sessions.Session) error {
	delete(session.Values, "items")
	return session.Save(c.Request, c.Writer)
}

// GetSession はセッションを外部から取得するためのヘルパー
func (h *CartHandler) GetSession(c *gin.Context) (*sessions.Session, cartData, error) {
	return func() (*sessions.Session, cartData, error) {
		data, session, err := h.getCart(c)
		return session, data, err
	}()
}
