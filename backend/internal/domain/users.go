package domain

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID        uint      `gorm:"primaryKey"`
	FullName  string    `gorm:"size:100;not null" json:"full_name"`
	Email     string    `gorm:"unique;not null"`
	Password  string    `gorm:"not null"`
	Role      string    `gorm:"default:'customer'"`
	Balance   float64   `gorm:"not null;default:0"`
	Pockets   []Pocket  `gorm:"foreignKey:UserID"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) CalculateBalance(db *gorm.DB) error {
	var totalBalance float64
	for _, pocket := range u.Pockets {
		totalBalance += pocket.Balance
	}
	u.Balance = totalBalance
	return nil
}
