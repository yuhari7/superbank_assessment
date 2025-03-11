package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yuhari7/superbank_assessment/internal/domain"
	"github.com/yuhari7/superbank_assessment/internal/infra"
)

func CreatePocketHandler(c *gin.Context) {
	var req struct {
		UserID        uint   `json:"user_id"`
		Name          string `json:"name"`
		AccountNumber string `json:"account_number"`
		BankName      string `json:"bank_name"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	pocket := domain.Pocket{
		UserID: req.UserID,
		Name:   req.Name,
		BankAccount: domain.BankAccount{
			AccountNumber: req.AccountNumber,
			BankName:      req.BankName,
		},
		Balance: 0, // Pocket selalu mulai dengan saldo 0
	}

	if err := infra.DB.Create(&pocket).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create pocket"})
		return
	}

	c.JSON(http.StatusCreated, pocket)
}
