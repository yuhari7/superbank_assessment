package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yuhari7/superbank_assessment/internal/middleware"
	"github.com/yuhari7/superbank_assessment/internal/usecase"
)

type CustomerHandler struct {
	usecase usecase.CustomerUsecase
}

func NewCustomerHandler(app *fiber.App, usecase usecase.CustomerUsecase) {
	handler := &CustomerHandler{usecase: usecase}

	// Secure route with JWT Middleware
	customerRoutes := app.Group("/customers", middleware.JWTMiddleware)
	customerRoutes.Get("/", handler.GetCustomers)
}

func (h *CustomerHandler) GetCustomers(c *fiber.Ctx) error {
	customers, err := h.usecase.FetchCustomers()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(customers)
}
