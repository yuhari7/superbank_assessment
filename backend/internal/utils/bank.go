package utils

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/yuhari7/superbank_assessment/internal/domain"
	"github.com/yuhari7/superbank_assessment/internal/infra"
)

// GenerateBankAccountNumber memastikan nomor unik
func GenerateBankAccountNumber() string {
	rand.Seed(time.Now().UnixNano())

	for {
		// Generate nomor 10 digit
		accountNumber := strconv.Itoa(1000000000 + rand.Intn(9000000000))

		// Cek apakah sudah ada di database
		var existingAccount domain.BankAccount
		result := infra.DB.Where("account_number = ?", accountNumber).First(&existingAccount)
		if result.RowsAffected == 0 {
			return accountNumber
		}
	}
}
