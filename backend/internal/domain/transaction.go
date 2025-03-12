package domain

import "time"

type Transaction struct {
	ID         uint      `gorm:"primaryKey"`
	SenderID   uint      `gorm:"not null"`
	ReceiverID uint      `gorm:"not null"`
	Amount     float64   `gorm:"not null"`
	CreatedAt  time.Time `json:"created_at"`
}
