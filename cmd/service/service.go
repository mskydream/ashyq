package service

import (
	"crypto/sha1"
	"fmt"

	"github.com/mskydream/qr-code/cmd/model"
	"github.com/mskydream/qr-code/cmd/repository"
)

type Authorization interface {
	CreateUser(user *model.User) (int, error)
}

type Service struct {
	repo *repository.Repository
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		repo: repos,
	}
}

func (s *Service) CreateUser(user *model.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	fmt.Println("BornDate in service: ", user.BornDate)
	// user.BornDate, err := time.Parse()
	return s.repo.CreateUser(user)
}

// генерация пароля
func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash)
}
