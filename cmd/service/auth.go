package service

import (
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/mskydream/qr-code/cmd/model"
)

const (
	salt       = "dksk232jko0239j9e029"
	signingKey = "sdksmdsomddm"
	tokenTTL   = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func (s *Service) CreateUser(user *model.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *Service) GenerateToken(iin, password string) (response model.GenerateTokenResponse, err error) {
	user, err := s.repo.GetUser(iin, generatePasswordHash(password))
	if err != nil {
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.ID,
	})

	signed, err := token.SignedString([]byte(signingKey))
	response.Token = signed
	return
}

// генерация пароля
func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
