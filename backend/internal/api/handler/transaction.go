package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yuhari7/superbank_assessment/internal/domain"
	"github.com/yuhari7/superbank_assessment/internal/infra"
)

func TransferHandler(c *gin.Context) {
	var req struct {
		FromPocketID uint    `json:"from_pocket_id"`
		ToUserID     uint    `json:"to_user_id"`
		Amount       float64 `json:"amount"`
	}

	// Bind JSON body to request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// 1. Cek apakah FromPocket ada di database
	var fromPocket domain.Pocket
	if err := infra.DB.First(&fromPocket, "id = ?", req.FromPocketID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "From Pocket not found"})
		return
	}

	// 2. Cek apakah saldo pengirim cukup
	if fromPocket.Balance < req.Amount {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Insufficient balance"})
		return
	}

	// 3. Cek apakah User penerima ada
	var toUser domain.User
	if err := infra.DB.First(&toUser, "id = ?", req.ToUserID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Recipient User not found"})
		return
	}

	// 4. Cek apakah penerima memiliki pocket (contohnya pocket utama)
	var toPocket domain.Pocket
	if err := infra.DB.First(&toPocket, "user_id = ?", req.ToUserID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Recipient's Pocket not found"})
		return
	}

	// 5. Lakukan transfer, kurangi saldo pengirim, tambah saldo penerima
	// Mulai transaksi
	tx := infra.DB.Begin()

	// Kurangi saldo dari Pocket pengirim
	fromPocket.Balance -= req.Amount
	if err := tx.Save(&fromPocket).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update sender's pocket"})
		return
	}

	// Tambah saldo ke Pocket penerima
	toPocket.Balance += req.Amount
	if err := tx.Save(&toPocket).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update recipient's pocket"})
		return
	}

	// Commit transaksi
	tx.Commit()

	// Kirim response sukses
	c.JSON(http.StatusOK, gin.H{"message": "Transfer successful"})
}
