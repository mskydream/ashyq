package service

import (
	"github.com/mskydream/ashyq/model"
	"github.com/mskydream/ashyq/repository"
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
	GetStatus(userId int, qrId string) (model.Status, error)
	GetVisits(userId int) ([]model.Visit, error)
	GetVisit(userId int, id string) (model.Visit, error)
}

type Service struct {
	Authorization
	RealEstate
	Visit
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		RealEstate:    NewRealEstateService(repos.RealEstate),
		Visit:         NewVisitService(repos.Visit),
	}
}
