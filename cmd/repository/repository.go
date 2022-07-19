package repository

import (
	"fmt"

	"github.com/mskydream/qr-code/cmd/db"
	"github.com/mskydream/qr-code/cmd/model"
)

type Authorization interface {
	CreateUser(user *model.User) (int, error)
	// GetUser(username, password string) (model.User, error)
}

type Repository struct {
	db *db.DB
}

func NewRepository(db *db.DB) *Repository {
	return &Repository{
		db: db,
	}
}

// func NewAuthRepository(db *sqlx.DB) *AuthPostgres {
// 	return &AuthPostgres{db: db}
// }
func (r *Repository) CreateUser(user *model.User) (id int, err error) {
	query := `INSERT INTO user_profile(name,surname, born, status,phone_number, iin,gender,residential_address,password,created_at)
								values($1, $2, $3, $4, $5, $6, $7, $8, $9, now()) RETURNING id`
	row := r.db.Conn.QueryRow(query, user.Name, user.Surname, user.BornDate, user.Status, user.PhoneNumber, user.IIN, user.Gender, user.ResidentalAddress, user.Password)
	fmt.Println("BornDate in repository", user.BornDate)
	if err = row.Scan(&id); err != nil {
		return
	}

	// _, err = r.db.Conn.NamedExec(`INSERT INTO user_profile (username,surname, born, status,phone_number, iin,gender,residental_address,password,created_at)VALUES (:username, :surname, :born, :status, phone_number, :iin, :gender, :residental_address, :password, NOW())`, user)
	// if err != nil {
	// 	fmt.Println("Error in repository")
	// }
	return
}

// func (r *AuthPostgres) GetUser(username, password string) (model.User, error) {
// 	var user model.User
// 	query := fmt.Sprintln("SELECT id FROM users WHERE username=$1 AND password_hash=$2")
// 	err := r.db.Get(&user, query, username, password)

// 	return user, err
// }
