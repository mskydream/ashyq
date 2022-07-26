package service

import (
	"github.com/mskydream/ashyq/api/model"
	"github.com/mskydream/ashyq/api/repository"
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
	CheckAddress(address *string) error
	Update(userId int, id string, realEstate *model.RealEstate) error
	Delete(userId int, id string) error
}
type Visit interface {
	CreateVisit(userId int, visit *model.Visit) (int, error)
	GetVisits(userId int) ([]model.Visit, error)
	GetVisit(userId int, id string) (model.Visit, error)
}

type Service struct {
	repo *repository.Repository
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		repo: repos,
	}
}
