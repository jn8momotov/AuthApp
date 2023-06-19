package service

import "AuthApp/internal/repository"

type AuthService struct {
	repo repository.Auth
}

func NewAuthService(repo repository.Auth) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) SignIn() error {
	return nil
}

func (s *AuthService) SignUp() error {
	return nil
}
