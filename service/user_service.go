package service

import (
	"errors"

	"github.com/EmmanoelDan/importador/model"
	"github.com/EmmanoelDan/importador/repository"
	"github.com/EmmanoelDan/importador/util"
)

type CreateUserService struct {
	CreateUserRepo *repository.UserRepository
}

func NewCreateUserService(createUserRepo *repository.UserRepository) *CreateUserService {
	return &CreateUserService{CreateUserRepo: createUserRepo}
}

func (s *CreateUserService) Register(usersername string, password string) (*model.User, error) {
	_, err := s.CreateUserRepo.FindByUsername(usersername)

	if err == nil {
		return nil, errors.New("User already registered")
	}

	hashedPassword, err := util.HashPassword(password)
	if err != nil {
		return nil, errors.New("Invalid password")
	}

	user := &model.User{
		Username: usersername,
		Password: hashedPassword,
	}

	err = s.CreateUserRepo.Create(user)

	if err != nil {
		return nil, errors.New("Could not create user")
	}

	return user, nil
}
