package repository

import (
	"github.com/mskydream/ashyq/cmd/model"
)

func (r *Repository) CreateVisit(userId int, visit *model.Visit) (int, error) {
	row := r.db.Conn.QueryRow("INSERT INTO visit (real_estate_id, user_profile_id, created_at) VALUES ($1, $2, NOW()) RETURNING id", visit.RealEstateId, userId)
	if err := row.Scan(&visit.Id); err != nil {
		return 0, err
	}
	return visit.Id, nil
}

func (r *Repository) GetVisits(userId int) (visits []model.Visit, err error) {
	return visits, r.db.Conn.Select(&visits, "SELECT id, real_estate_id, user_profile_id, created_at FROM visit WHERE user_profile_id = $1", userId)
}

func (r *Repository) GetVisit(userId int, id string) (visit model.Visit, err error) {
	return visit, r.db.Conn.Get(&visit, "SELECT id, real_estate_id, user_profile_id, created_at FROM visit WHERE user_profile_id = $1 AND id = $2", userId, id)
}
