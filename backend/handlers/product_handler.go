package handlers

import (
	"net/http"
	"strconv"

	"github.com/example/go-shopping/services"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	svc *services.ProductService
}

func NewProductHandler(svc *services.ProductService) *ProductHandler {
	return &ProductHandler{svc: svc}
}

func (h *ProductHandler) GetProducts(c *gin.Context) {
	products, err := h.svc.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}

func (h *ProductHandler) SearchProducts(c *gin.Context) {
	keyword := c.Query("keyword")
	minPriceStr := c.Query("minPrice")
	maxPriceStr := c.Query("maxPrice")

	var keywordPtr *string
	if keyword != "" {
		keywordPtr = &keyword
	}

	var minPrice, maxPrice *float64
	if minPriceStr != "" {
		v, err := strconv.ParseFloat(minPriceStr, 64)
		if err != nil || v < 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid minPrice"})
			return
		}
		minPrice = &v
	}
	if maxPriceStr != "" {
		v, err := strconv.ParseFloat(maxPriceStr, 64)
		if err != nil || v < 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid maxPrice"})
			return
		}
		maxPrice = &v
	}
	if minPrice != nil && maxPrice != nil && *minPrice > *maxPrice {
		c.JSON(http.StatusBadRequest, gin.H{"error": "minPrice must be <= maxPrice"})
		return
	}

	products, err := h.svc.SearchProducts(keywordPtr, minPrice, maxPrice)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}

func (h *ProductHandler) GetProductsByCategory(c *gin.Context) {
	category := c.Param("category")
	products, err := h.svc.GetProductsByCategory(category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}

func (h *ProductHandler) FilterLow(c *gin.Context) {
	max := services.LowPriceMax
	products, err := h.svc.SearchProducts(nil, nil, &max)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}

func (h *ProductHandler) FilterMid(c *gin.Context) {
	min := services.MidPriceMin
	max := services.MidPriceMax
	products, err := h.svc.SearchProducts(nil, &min, &max)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}

func (h *ProductHandler) FilterHigh(c *gin.Context) {
	min := services.HighPriceMin
	products, err := h.svc.SearchProducts(nil, &min, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}
