package usecase

import (
	"errors"

	"github.com/yuhari7/superbank_assessment/pkg"

	"github.com/yuhari7/superbank_assessment/internal/entity"
	"github.com/yuhari7/superbank_assessment/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase interface {
	Register(username, password string) error
	Login(username, password string) (string, error)
}

type userUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) UserUsecase {
	return &userUsecase{userRepo: repo}
}

func (u *userUsecase) Register(username, password string) error {
	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &entity.User{Username: username, Password: string(hashedPassword)}
	return u.userRepo.CreateUser(user)
}

func (u *userUsecase) Login(username, password string) (string, error) {
	user, err := u.userRepo.GetUserByUsername(username)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	// Compare passwords
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid credentials")
	}

	// Generate JWT token
	token, err := pkg.GenerateJWT(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}
