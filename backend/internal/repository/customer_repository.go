package repository

import (
	"github.com/yuhari7/superbank_assessment/internal/entity"
	"gorm.io/gorm"
)

type CustomerRepository interface {
	GetAllCustomers() ([]entity.Customer, error)
}

type customerRepo struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return &customerRepo{db: db}
}

func (r *customerRepo) GetAllCustomers() ([]entity.Customer, error) {
	var customers []entity.Customer
	if err := r.db.Find(&customers).Error; err != nil {
		return nil, err
	}
	return customers, nil
}
