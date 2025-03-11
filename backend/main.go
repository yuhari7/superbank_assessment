package main

import (
	"github.com/gin-gonic/gin"
	handlers "github.com/yuhari7/superbank_assessment/internal/api/handler"
	"github.com/yuhari7/superbank_assessment/internal/infra"
	"github.com/yuhari7/superbank_assessment/internal/middleware"
	"github.com/yuhari7/superbank_assessment/migrations"
)

func main() {
	// Menghubungkan ke database
	infra.ConnectDB()

	// Melakukan migrasi database
	migrations.Migrate()

	// Membuat router Gin
	router := gin.Default()

	// Endpoint untuk login
	router.POST("/login", handlers.LoginHandler)

	// Protected route untuk transfer (gunakan AuthMiddleware)
	protected := router.Group("/protected")
	protected.Use(middleware.AuthMiddleware())
	protected.POST("/transfer", handlers.TransferHandler)

	// Menjalankan server
	router.Run(":8080")
}
