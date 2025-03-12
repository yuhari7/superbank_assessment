package migrations

import (
	"log"

	"github.com/yuhari7/superbank_assessment/internal/domain"
	"github.com/yuhari7/superbank_assessment/internal/infra"
)

func Migrate() {
	db := infra.DB

	err := db.Migrator().DropTable(&domain.User{}, &domain.Pocket{}, &domain.Transaction{})
	if err != nil {
		log.Fatal("Failed to drop tables: ", err)
	}

	err = db.AutoMigrate(
		&domain.User{},
		&domain.Pocket{},
		&domain.TermDeposit{},
		&domain.Transaction{},
	)
	if err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}

	// 🔹 Seed Initial Users
	seedUsers()
}

func seedUsers() {
	db := infra.DB
	users := []domain.User{
		{
			FullName: "Admin Fintech",
			Email:    "admin@fintech.com",
			Password: "admin123",
			Role:     "admin",
			Pockets: []domain.Pocket{
				{Name: "Main Pocket", Balance: 10000, BankAccount: domain.BankAccount{AccountNumber: "123456789", BankName: "FinTech Bank"}},
			},
		},
		{
			FullName: "User Fintech",
			Email:    "user@fintech.com",
			Password: "user123",
			Role:     "customer",
			Pockets: []domain.Pocket{
				{Name: "Main Pocket", Balance: 5000, BankAccount: domain.BankAccount{AccountNumber: "987654321", BankName: "FinTech Bank"}},
			},
		},
	}

	for _, user := range users {
		if err := user.HashPassword(); err != nil {
			log.Fatal("Failed to hash password: ", err)
		}
		db.Create(&user)
		user.CalculateBalance(db)
		db.Save(&user)
	}
}
