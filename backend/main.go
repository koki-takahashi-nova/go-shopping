package main

import (
	"log"
	"net/http"
	"os"

	"github.com/example/go-shopping/config"
	"github.com/example/go-shopping/handlers"
	"github.com/example/go-shopping/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

func main() {
	db, err := config.InitDB()
	if err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}

	sessionSecret := os.Getenv("SESSION_SECRET")
	if sessionSecret == "" {
		sessionSecret = "dev-secret-key-change-in-production"
	}
	store := sessions.NewCookieStore([]byte(sessionSecret))
	store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}

	productSvc := services.NewProductService(db)
	orderSvc := services.NewOrderService(db)

	productHandler := handlers.NewProductHandler(productSvc)
	cartHandler := handlers.NewCartHandler(store, db)
	orderHandler := handlers.NewOrderHandler(orderSvc, cartHandler)

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Accept"},
		AllowCredentials: true,
	}))

	api := r.Group("/api")
	{
		products := api.Group("/products")
		{
			products.GET("", productHandler.GetProducts)
			products.GET("/search", productHandler.SearchProducts)
			products.GET("/category/:category", productHandler.GetProductsByCategory)
			products.GET("/filter/low", productHandler.FilterLow)
			products.GET("/filter/mid", productHandler.FilterMid)
			products.GET("/filter/high", productHandler.FilterHigh)
		}

		cart := api.Group("/cart")
		{
			cart.GET("", cartHandler.GetCart)
			cart.POST("/add/:id", cartHandler.AddToCart)
			cart.DELETE("/remove/:id", cartHandler.RemoveFromCart)
		}

		api.POST("/orders", orderHandler.PlaceOrder)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server starting on :%s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
