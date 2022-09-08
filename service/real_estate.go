package service

import (
	"github.com/mskydream/ashyq/model"
	"github.com/mskydream/ashyq/repository"
)

type RealEstateService struct {
	repo repository.RealEstate
}

func NewRealEstateService(repo repository.RealEstate) *RealEstateService {
	return &RealEstateService{repo: repo}
}

func (s *RealEstateService) Create(userId int, realEstate *model.RealEstate) (int, error) {
	return s.repo.Create(userId, realEstate)
}

func (s *RealEstateService) GetAll(userId int) ([]model.RealEstate, error) {
	return s.repo.GetAll(userId)
}

func (s *RealEstateService) Get(userId int, id string) (model.RealEstate, error) {
	return s.repo.Get(userId, id)
}

func (s *RealEstateService) CheckAddress(address *string) error {
	return s.repo.CheckAddress(address)
}

func (s *RealEstateService) Delete(userId int, id string) error {
	return s.repo.Delete(userId, id)
}

func (s *RealEstateService) Update(userId int, id string, realEstate *model.RealEstate) error {
	return s.repo.Update(userId, id, realEstate)
}
