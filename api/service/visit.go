package service

import "github.com/mskydream/ashyq/api/model"

func (s *Service) CreateVisit(userId int, qr *model.Qr) (string, error) {
	return s.repo.CreateVisit(userId, qr)
}

func (s *Service) GetVisits(userId int) ([]model.Visit, error) {
	return s.repo.GetVisits(userId)
}

func (s *Service) GetVisit(userId int, id string) (model.Visit, error) {
	return s.repo.GetVisit(userId, id)
}
