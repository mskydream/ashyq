package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/mskydream/ashyq/model"
)

type VisitPostgres struct {
	db *sqlx.DB
}

func NewVisitPostgres(db *sqlx.DB) *VisitPostgres {
	return &VisitPostgres{db: db}
}

func (r *VisitPostgres) GetStatus(userId int, qrId string) (model.Status, error) {
	var realEstateId string
	err := r.db.Get(&realEstateId, "SELECT id FROM real_estate WHERE qr_code = $1", qrId)
	if err != nil {
		return model.Status{}, err
	}

	_, err = r.db.Exec("INSERT INTO visit (real_estate_id, user_profile_id, created_at) VALUES ($1, $2, NOW())", realEstateId, userId)
	if err != nil {
		return model.Status{}, err
	}

	var status string
	err = r.db.Get(&status, "SELECT status FROM user_profile WHERE id = $1", userId)
	if err != nil {
		return model.Status{}, err
	}
	return model.Status{Status: status}, nil
}

func (r *VisitPostgres) GetVisits(userId int) (visits []model.Visit, err error) {
	return visits, r.db.Select(&visits, "SELECT id, real_estate_id, user_profile_id, created_at FROM visit WHERE user_profile_id = $1", userId)
}

func (r *VisitPostgres) GetVisit(userId int, id string) (visit model.Visit, err error) {
	return visit, r.db.Get(&visit, "SELECT id, real_estate_id, user_profile_id, created_at FROM visit WHERE user_profile_id = $1 AND id = $2", userId, id)
}

func (r *VisitPostgres) CheckQr(address string) (realEstateId string, err error) {
	return realEstateId, r.db.Get(&realEstateId, "SELECT id FROM real_estate WHERE qr_code = $1", address)
}
