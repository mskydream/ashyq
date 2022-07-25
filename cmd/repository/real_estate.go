package repository

import (
	"github.com/mskydream/ashyq/cmd/model"
	"github.com/skip2/go-qrcode"
)

func (r *Repository) Create(userId int, realEstate *model.RealEstate) (id int, err error) {
	tx, err := r.db.Conn.Begin()
	if err != nil {
		return 0, err
	}

	err = tx.QueryRow("INSERT INTO real_estate (user_profile_id,  address, qr_code,created_at) VALUES ($1, $2, $3, NOW()) RETURNING id", userId, realEstate.Address, realEstate.QrCode).Scan(&id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (r *Repository) GetAll(userId int) (realEstates []model.RealEstate, err error) {
	rows, err := r.db.Conn.Query("SELECT id, address, qr_code, created_at FROM real_estate WHERE user_profile_id = $1", userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var realEstate model.RealEstate
		err = rows.Scan(&realEstate.Id, &realEstate.Address, &realEstate.QrCode, &realEstate.CreatedAt)
		if err != nil {
			return nil, err
		}
		realEstates = append(realEstates, realEstate)
	}
	return realEstates, nil
}

func (r *Repository) Get(userId int, id string) (realEstate model.RealEstate, err error) {
	err = r.db.Conn.QueryRow("SELECT id, address, qr_code, created_at FROM real_estate WHERE user_profile_id = $1 AND id = $2", userId, id).Scan(&realEstate.Id, &realEstate.Address, &realEstate.QrCode, &realEstate.CreatedAt)
	if err != nil {
		return model.RealEstate{}, err
	}
	return realEstate, nil
}

func WriteQRCodeToFile(pathFileName, data string) error {
	return qrcode.WriteFile(data, qrcode.Medium, 256, pathFileName)
}

func (r *Repository) CheckAddress(address string) error {
	return r.db.Conn.Select("SELECT * FROM real_estate WHERE address = $1", address)
}
