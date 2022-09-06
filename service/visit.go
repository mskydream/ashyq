package service

import "github.com/mskydream/ashyq/model"

func (s *Service) GetStatus(userId int, qrId string) (model.Status, error) {
	return s.repo.GetStatus(userId, qrId)
}

func (s *Service) GetVisits(userId int) ([]model.Visit, error) {
	return s.repo.GetVisits(userId)
}

func (s *Service) GetVisit(userId int, id string) (model.Visit, error) {
	return s.repo.GetVisit(userId, id)
}
