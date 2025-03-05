package service

import (
	"github.com/EmmanoelDan/importador/repository"
	"github.com/EmmanoelDan/importador/util"
)

type AuthService struct {
	UserRepo repository.UserRepository
}

func NewAuthService(userRepo repository.UserRepository) *AuthService {
	return &AuthService{UserRepo: userRepo}
}

func (s *AuthService) Authenticate(username, password string) (string, error) {
	user, err := s.UserRepo.FindByUsername(username)
	if err != nil {
		return "", err
	}

	if !util.ComparePassword(user.Password, password) {
		return "", err
	}

	token, err := util.GenerateJWT(username)
	if err != nil {
		return "", err
	}

	return token, nil

}
