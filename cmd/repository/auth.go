package repository

import (
	"github.com/mskydream/qr-code/cmd/model"
)

func (r *Repository) CreateUser(user *model.User) (id int, err error) {
	query := `INSERT INTO user_profile(name,surname, born, status,phone_number, iin,gender,residential_address,password,created_at)
								values($1, $2, $3, $4, $5, $6, $7, $8, $9, now()) RETURNING id`
	row := r.db.Conn.QueryRow(query, user.Name, user.Surname, user.BornDate, user.Status, user.PhoneNumber, user.IIN, user.Gender, user.ResidentialAddress, user.Password)

	if err = row.Scan(&id); err != nil {
		return
	}
	return
}

func (r *Repository) GetUser(iin, password string) (user model.User, err error) {
	err = r.db.Conn.Get(&user, "SELECT id FROM user_profile WHERE iin=$1 AND password=$2", iin, password)
	return
}
