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

// func (s *Service) generateQrCode(address string) (string, error) {
// 	nameQrCode := time.Now().Unix()
// 	// nameQrCode to string
// 	nameQrCodeStr := strconv.FormatInt(nameQrCode, 10) + ".png"
// 	err := qrcode.WriteFile("http://localhost:8080/qr-code/"+address+".png", qrcode.Medium, 256, nameQrCodeStr) // create a QR code image
// 	if err != nil {
// 		return "", err
// 	}
// 	return address, nil
// }
