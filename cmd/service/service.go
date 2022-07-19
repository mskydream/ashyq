package service

import (
	"github.com/mskydream/qr-code/cmd/model"
	"github.com/mskydream/qr-code/cmd/repository"
)

type Authorization interface {
	CreateUser(user *model.User) (int, error)
	GenerateToken(username, password string) (model.GenerateTokenResponse, error)
}

type Service struct {
	repo *repository.Repository
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		repo: repos,
	}
}
