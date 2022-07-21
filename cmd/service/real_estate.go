package service

import (
	"github.com/mskydream/qr-code/cmd/model"
)

func (s *Service) Create(userId int, realEstate *model.RealEstate) (int, error) {
	return s.repo.Create(userId, realEstate)
}

func (s *Service) GetAll(userId int) ([]model.RealEstate, error) {
	return s.repo.GetAll(userId)
}

func (s *Service) Get(userId int, id string) (model.RealEstate, error) {
	return s.repo.Get(userId, id)
}
