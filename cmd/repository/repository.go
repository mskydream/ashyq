package repository

import (
	"github.com/mskydream/qr-code/cmd/db"
	"github.com/mskydream/qr-code/cmd/model"
)

type Authorization interface {
	CreateUser(user *model.User) (int, error)
	GetUser(iin, password string) (model.User, error)
}

type Repository struct {
	db *db.DB
}

func NewRepository(db *db.DB) *Repository {
	return &Repository{
		db: db,
	}
}
