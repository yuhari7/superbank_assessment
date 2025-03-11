package domain

import "time"

type TermDeposit struct {
	ID        uint    `gorm:"primaryKey"`
	UserID    uint    `gorm:"not null"`
	Amount    float64 `gorm:"not null"`
	Duration  int     `gorm:"not null"` // Duration in months
	Interest  float64 `gorm:"not null"` // Interest rate
	CreatedAt time.Time
	UpdatedAt time.Time
}
