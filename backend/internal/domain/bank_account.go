package domain

type BankAccount struct {
	AccountNumber string `gorm:"unique;not null"`
	BankName      string `gorm:"not null"`
	Balance       float64
}
