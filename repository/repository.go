package repository

import (
	"github.com/mskydream/ashyq/db"
	"github.com/mskydream/ashyq/model"
)

type Authorization interface {
	CreateUser(user *model.User) (int, error)
	GetUser(iin, password string) (model.User, error)
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

type Repository struct {
	db *db.DB
}

func NewRepository(db *db.DB) *Repository {
	return &Repository{
		db: db,
	}
}
