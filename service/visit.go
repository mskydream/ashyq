package service

import (
	"github.com/mskydream/ashyq/model"
	"github.com/mskydream/ashyq/repository"
)

type VisitService struct {
	repo repository.Visit
}

func NewVisitService(repo repository.Visit) *VisitService {
	return &VisitService{repo: repo}
}

func (s *VisitService) GetStatus(userId int, qrId string) (model.Status, error) {
	return s.repo.GetStatus(userId, qrId)
}

func (s *VisitService) GetVisits(userId int) ([]model.Visit, error) {
	return s.repo.GetVisits(userId)
}

func (s *VisitService) GetVisit(userId int, id string) (model.Visit, error) {
	return s.repo.GetVisit(userId, id)
}
