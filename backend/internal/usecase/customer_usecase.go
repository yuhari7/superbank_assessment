package usecase

import (
	"github.com/yuhari7/superbank_assessment/internal/entity"
	"github.com/yuhari7/superbank_assessment/internal/repository"
)

type CustomerUsecase interface {
	FetchCustomers() ([]entity.Customer, error)
}

type customerUsecase struct {
	customerRepo repository.CustomerRepository
}

func NewCustomerUsecase(repo repository.CustomerRepository) CustomerUsecase {
	return &customerUsecase{customerRepo: repo}
}

func (u *customerUsecase) FetchCustomers() ([]entity.Customer, error) {
	return u.customerRepo.GetAllCustomers()
}
