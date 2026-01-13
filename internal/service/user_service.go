package service

import (
	"fmt"
	"go-web/internal/entity"
	"go-web/internal/repository"
	"go-web/pkg/security"
)

type UserService struct {
	repository repository.UserRepository
}

func NewUserService() IUserService {
	return &UserService{
		repository: *repository.NewUserRepository(),
	}
}

func (s *UserService) CreateUser(login string, password string) {
	passwordHash, err := security.HashPassword(password)
	if err != nil {
		fmt.Println("An error occured while hashing password.")
	}

	entityUser := entity.User{
		Login:    login,
		Password: passwordHash,
	}

	repository := repository.NewUserRepository()
	err = repository.Insert(entityUser)
	if err != nil {
		fmt.Println(err.Error())
	}
}
