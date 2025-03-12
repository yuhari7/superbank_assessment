package main

import (
	"github.com/gin-gonic/gin"
	handlers "github.com/yuhari7/superbank_assessment/internal/api/handler"
	"github.com/yuhari7/superbank_assessment/internal/infra"
	"github.com/yuhari7/superbank_assessment/internal/middleware"
	"github.com/yuhari7/superbank_assessment/migrations"
)

func main() {
	// Connect to the database
	infra.ConnectDB()

	// Run migrations
	migrations.Migrate()

	// Create router
	router := gin.Default()

	// Login endpoint
	router.POST("/login", handlers.LoginHandler)

	// Protected routes
	protected := router.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	protected.POST("/transfer", handlers.TransferHandler)
	protected.POST("/pocket", handlers.CreatePocketHandler)
	protected.GET("/balance", handlers.GetBalanceHandler)
	protected.GET("/term-deposits", handlers.GetTermDepositsHandler)

	// Run the server
	router.Run(":8080")
}
