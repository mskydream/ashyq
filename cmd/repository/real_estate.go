package repository

import "github.com/mskydream/qr-code/cmd/model"

// func Create(model model.RealEstate) (int error)      {
// 	query := "INSERT INTO real_estate(user_profile_id)"
// }
func (r *Repository) GetRealEstate(id int) (realEstate model.RealEstate, err error) {
	err = r.db.Conn.Get(&realEstate, "SELECT * FROM user_profile WHERE id=$1", id)
	return
}
