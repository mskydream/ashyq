package model

import "time"

type Response struct {
	IsSuccess bool        `json:"isSuccess"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
}

type User struct {
	Id                 int       `db:"id" json:"id"`
	Name               string    `db:"name" json:"name"`
	Surname            string    `db:"surname" json:"surname"`
	BornDate           time.Time `db:"born_date" json:"born_date"`
	Status             string    `db:"status" json:"status"`
	PhoneNumber        string    `db:"phone_number" json:"phone_number"`
	IIN                string    `db:"iin" json:"iin"`
	Gender             string    `db:"gender" json:"gender"`
	ResidentialAddress string    `db:"residential_address" json:"residential_address"`
	Password           string    `db:"password" json:"password"`
	CreatedAt          time.Time `db:"created_at" json:"created_at"`
}

type RealEstate struct {
	Id            uint      `db:"id" json:"id"`
	UserProfileId uint      `db:"user_profile_id" json:"user_profile_id"`
	Address       string    `db:"address" json:"address"`
	QrCode        string    `db:"qr_code" json:"qr_code"`
	CreatedAt     time.Time `db:"created_at" json:"created_at"`
}

type Visit struct {
	Id            int       `db:"id" json:"id"`
	RealEstateId  string    `db:"real_estate_id" json:"real_estate_id"`
	UserProfileId string    `db:"user_profile_id" json:"user_profile_id"`
	CreatedAt     time.Time `db:"created_at" json:"created_at"`
}

type SignInInput struct {
	IIN      string `json:"iin" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type GenerateTokenResponse struct {
	Token string `json:"token"`
}

type Status struct {
	Status string `json:"status"`
}
