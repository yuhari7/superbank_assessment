package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yuhari7/superbank_assessment/internal/domain"
	"github.com/yuhari7/superbank_assessment/internal/infra"
)

func CreateTermDepositHandler(c *gin.Context) {
	var req struct {
		UserID   uint    `json:"user_id"`
		Amount   float64 `json:"amount"`
		Duration int     `json:"duration"`
		Interest float64 `json:"interest"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	termDeposit := domain.TermDeposit{
		UserID:   req.UserID,
		Amount:   req.Amount,
		Duration: req.Duration,
		Interest: req.Interest,
	}

	if err := infra.DB.Create(&termDeposit).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create term deposit"})
		return
	}

	c.JSON(http.StatusCreated, termDeposit)
}

func GetTermDepositsHandler(c *gin.Context) {
	var req struct {
		UserID uint `json:"user_id"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	var termDeposits []domain.TermDeposit
	if err := infra.DB.Where("user_id = ?", req.UserID).Find(&termDeposits).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Term deposits not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"term_deposits": termDeposits})
}
