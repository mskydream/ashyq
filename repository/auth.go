package repository

import (
	"errors"
	"strconv"

	"github.com/mskydream/ashyq/model"
)

func (r *Repository) CreateUser(user *model.User) (id int, err error) {
	if !checkIIN(user.IIN) || !checkPhoneNumber(user.PhoneNumber) {
		return 0, errors.New("invalid iin or phone number")
	}

	query := `INSERT INTO user_profile(name,surname, born, status,phone_number, iin,gender,residential_address,password,created_at)
								values($1, $2, $3, $4, $5, $6, $7, $8, $9, now()) RETURNING id`
	row := r.db.Conn.QueryRow(query, user.Name, user.Surname, user.BornDate, user.Status, user.PhoneNumber, user.IIN, user.Gender, user.ResidentialAddress, user.Password)

	if err = row.Scan(&id); err != nil {
		return
	}
	return
}

func (r *Repository) GetUser(iin, password string) (user model.User, err error) {
	return user, r.db.Conn.Get(&user, "SELECT id FROM user_profile WHERE iin=$1 AND password=$2", iin, password)
}

func checkIIN(iin string) bool {
	if len(iin) == 12 {
		for i := 0; i < len(iin); i++ {
			_, err := strconv.Atoi(string(iin[i]))
			if err != nil {
				return false
			}
		}
		return true
	}
	return false
}

func checkPhoneNumber(number string) bool {
	if len(number) == 11 {
		if string(number[0]) == "8" && string(number[1]) == "7" {
			for i := 2; i < len(number); i++ {
				_, err := strconv.Atoi(string(number[i]))
				if err != nil {
					return false
				}
			}
			return true
		}
		return false
	}
	return false
}
