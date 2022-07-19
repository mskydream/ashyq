package service

import "github.com/mskydream/qr-code/cmd/model"

func (s *Service) GetRealEstate(id int) (realEstate model.RealEstate, err error) {
	return s.repo.GetRealEstate(id)
}
