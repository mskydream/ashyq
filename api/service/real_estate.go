package service

import "github.com/mskydream/ashyq/api/model"

func (s *Service) Create(userId int, realEstate *model.RealEstate) (int, error) {
	return s.repo.Create(userId, realEstate)
}

func (s *Service) GetAll(userId int) ([]model.RealEstate, error) {
	return s.repo.GetAll(userId)
}

func (s *Service) Get(userId int, id string) (model.RealEstate, error) {
	return s.repo.Get(userId, id)
}

func (s *Service) CheckAddress(address string) error {
	return s.repo.CheckAddress(address)
}

func (s *Service) Delete(userId int, id string) error {
	return s.repo.Delete(userId, id)
}

func (s *Service) Update(userId int, id string, realEstate *model.RealEstate) error {
	return s.repo.Update(userId, id, realEstate)
}
