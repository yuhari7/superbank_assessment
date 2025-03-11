package migrations

import (
	"log"

	"github.com/yuhari7/superbank_assessment/internal/domain"
	"github.com/yuhari7/superbank_assessment/internal/infra"
)

func Migrate() {
	db := infra.DB

	// Auto Migrate Model
	err := db.AutoMigrate(&domain.User{})
	if err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}

	// Seed User Awal
	seedUsers()
}

func seedUsers() {
	db := infra.DB
	users := []domain.User{
		{FullName: "Admin Fintech", Email: "admin@fintech.com", Password: "admin123", Role: "admin"},
		{FullName: "User Fintech", Email: "user@fintech.com", Password: "user123", Role: "customer"},
	}

	for _, user := range users {
		if err := user.HashPassword(); err != nil {
			log.Fatal("Failed to hash password: ", err)
		}
		db.Create(&user)
	}
}
