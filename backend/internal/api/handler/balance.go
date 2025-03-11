package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yuhari7/superbank_assessment/internal/domain"
	"github.com/yuhari7/superbank_assessment/internal/infra"
)

func GetBalanceHandler(c *gin.Context) {
	var req struct {
		UserID uint `json:"user_id"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	var user domain.User
	if err := infra.DB.First(&user, req.UserID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"balance": user.Balance})
}
