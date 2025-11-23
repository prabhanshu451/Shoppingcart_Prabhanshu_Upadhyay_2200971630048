package main

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Enable CORS before any routes
	r.Use(CORSMiddleware())

	// ======== USER ROUTES ========
	r.POST("/users", CreateUser)
	r.GET("/users", ListUsers)
	r.POST("/users/login", LoginUser)

	// ======== PUBLIC LISTING ROUTES (NO AUTH REQUIRED) ========
	r.GET("/items", ListItems)
	r.GET("/carts", ListCarts)
	r.GET("/orders", ListOrders)

	// ======== AUTH REQUIRED ROUTES ========
	auth := r.Group("/")
	auth.Use(AuthMiddleware())

	auth.POST("/carts", CreateOrAddToCart)
	auth.POST("/orders", CreateOrderFromCart)

	return r
}
