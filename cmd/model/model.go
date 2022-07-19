package model

import "time"

type Response struct {
	IsSuccess bool        `json:"isSuccess"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
}

type User struct {
	ID                int       `db:"id" json:"id"`
	Name              string    `db:"name" json:"name"`
	Surname           string    `db:"surname" json:"surname"`
	BornDate          time.Time `db:"born_date" json:"born_date"`
	Status            string    `db:"status" json:"status"`
	PhoneNumber       string    `db:"phone_number" json:"phone_number"`
	IIN               string    `db:"iin" json:"iin"`
	Gender            string    `db:"gender" json:"gender"`
	ResidentalAddress string    `db:"residential_address" json:"residential_address"`
	Password          string    `db:"password" json:"password"`
	CreatedAt         time.Time `db:"created_at" json:"created_at"`
}

type RealEstate struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_profile_id"`
	Address   string `json:"address"`
	QrCode    string `json:"qr_code"`
	CreatedAt string `json:"created_at"`
}

type Visit struct {
	ID           int    `json:"id"`
	RealEstateID int    `json:"real_estate_id"`
	UserID       int    `json:"user_profile_id"`
	CreatedAt    string `json:"created_at"`
}

type SignInInput struct {
	IIN      string `json:"iin" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type GenerateTokenResponse struct {
	Token string `json:"token"`
}
