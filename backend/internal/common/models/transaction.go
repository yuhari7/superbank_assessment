package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	SenderID   uint    `gorm:"not null"`
	ReceiverID uint    `gorm:"not null"`
	Amount     float64 `gorm:"not null"`
	Status     string  `gorm:"type:varchar(20);default:'pending'"` // pending, completed, failed
}
