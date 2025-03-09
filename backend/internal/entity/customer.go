package entity

type Customer struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `json:"name"`
}
