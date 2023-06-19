package service

import "AuthApp/internal/repository"

type Auth interface {
	SignIn() error
	SignUp() error
}

type Services struct {
	Auth Auth
}

func NewServices(repos *repository.Repositories) *Services {
	return &Services{
		Auth: NewAuthService(repos.Auth),
	}
}
