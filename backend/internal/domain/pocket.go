package domain

import "time"

type Pocket struct {
	ID          uint        `gorm:"primaryKey"`
	UserID      uint        `gorm:"not null"`
	Name        string      `gorm:"not null"`
	Balance     float64     `gorm:"default:0"`
	BankAccount BankAccount `gorm:"embedded"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
