package main

import (
	"github.com/gin-gonic/gin"
	handlers "github.com/yuhari7/superbank_assessment/internal/api/handler"
	"github.com/yuhari7/superbank_assessment/internal/infra"
	"github.com/yuhari7/superbank_assessment/internal/middleware"
	"github.com/yuhari7/superbank_assessment/migrations"
)

func main() {
	infra.ConnectDB()
	migrations.Migrate()

	router := gin.Default()

	// Endpoint authentication
	router.POST("/login", handlers.LoginHandler)

	// Protected route
	protected := router.Group("/protected")
	protected.Use(middleware.AuthMiddleware())
	protected.GET("/data", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Protected content"})
	})

	router.Run(":8080")
}
