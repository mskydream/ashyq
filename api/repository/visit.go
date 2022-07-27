package repository

import (
	"github.com/mskydream/ashyq/api/model"
)

func (r *Repository) CreateVisit(userId int, qr *model.Qr) (string, error) {
	realEstateId, err := r.CheckQr(qr.Qr)
	if err != nil {
		return "", err
	}
	// return int and error
	tx, err := r.db.Conn.Begin()
	if err != nil {
		return "", err
	}

	defer tx.Rollback()

	_, err = tx.Exec("INSERT INTO visit (real_estate_id, user_profile_id, created_at) VALUES ($1, $2, NOW())", realEstateId, userId)
	if err != nil {
		return "", err
	}

	status, err := r.GetStatus(userId)
	if err != nil {
		return "", err
	}

	return status, tx.Commit()
}

func (r *Repository) GetVisits(userId int) (visits []model.Visit, err error) {
	return visits, r.db.Conn.Select(&visits, "SELECT id, real_estate_id, user_profile_id, created_at FROM visit WHERE user_profile_id = $1", userId)
}

func (r *Repository) GetVisit(userId int, id string) (visit model.Visit, err error) {
	return visit, r.db.Conn.Get(&visit, "SELECT id, real_estate_id, user_profile_id, created_at FROM visit WHERE user_profile_id = $1 AND id = $2", userId, id)
}

func (r *Repository) CheckQr(address string) (realEstateId string, err error) {
	return realEstateId, r.db.Conn.Get(&realEstateId, "SELECT id FROM real_estate WHERE qr_code = $1", address)
}

func (r *Repository) GetStatus(userId int) (status string, err error) {
	return status, r.db.Conn.Get(&status, "SELECT status FROM user_profile WHERE id = $1", userId)
}
