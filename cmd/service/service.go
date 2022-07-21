package service

import (
	"github.com/mskydream/qr-code/cmd/model"
	"github.com/mskydream/qr-code/cmd/repository"
)

type Authorization interface {
	CreateUser(user *model.User) (int, error)
	GenerateToken(username, password string) (model.GenerateTokenResponse, error)
	ParseToken(token string) (int, error)
}

type RealEstate interface {
	Create(userId int, realEstate *model.RealEstate) (int, error)
	GetAll(userId int) ([]model.RealEstate, error)
	Get(userId int, id string) (model.RealEstate, error)
	// Update(userId int, id int, realEstate *model.RealEstate) error
	// Delete(userId int, id int) error
}

type Service struct {
	repo *repository.Repository
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		repo: repos,
	}
}
