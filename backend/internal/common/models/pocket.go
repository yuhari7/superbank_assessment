package models

import "gorm.io/gorm"

type Pocket struct {
	gorm.Model
	UserID  uint    `gorm:"not null"`
	Name    string  `gorm:"type:varchar(100);not null"`
	Balance float64 `gorm:"default:0"`
}
