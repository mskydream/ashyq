package repository

import (
	"github.com/mskydream/qr-code/cmd/db"
	"github.com/mskydream/qr-code/cmd/model"
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
	// Update(userId int, id int, realEstate *model.RealEstate) error
	// Delete(userId int, id int) error
}

type Repository struct {
	db *db.DB
}

func NewRepository(db *db.DB) *Repository {
	return &Repository{
		db: db,
	}
}
