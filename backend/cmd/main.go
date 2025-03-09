package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/yuhari7/superbank_assessment/internal/entity"
	"github.com/yuhari7/superbank_assessment/internal/handler"
	"github.com/yuhari7/superbank_assessment/internal/repository"
	"github.com/yuhari7/superbank_assessment/internal/usecase"
	"github.com/yuhari7/superbank_assessment/pkg"
)

func main() {
	app := fiber.New()
	db := pkg.ConnectDB()

	// Migrate Tables
	db.AutoMigrate(&entity.User{}, &entity.Customer{})

	// Dependency Injection
	userRepo := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo)
	handler.NewUserHandler(app, userUsecase)

	customerRepo := repository.NewCustomerRepository(db)
	customerUsecase := usecase.NewCustomerUsecase(customerRepo)
	handler.NewCustomerHandler(app, customerUsecase)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	app.Listen(":" + port)
}
