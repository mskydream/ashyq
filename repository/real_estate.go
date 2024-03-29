package repository

import (
	"os"
	"strconv"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/mskydream/ashyq/model"
	"github.com/skip2/go-qrcode"
)

type RealEstatePostgres struct {
	db *sqlx.DB
}

func NewRealEstatePostgres(db *sqlx.DB) *RealEstatePostgres {
	return &RealEstatePostgres{db: db}
}

func (r *RealEstatePostgres) Create(userId int, realEstate *model.RealEstate) (id int, err error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	realEstate.QrCode, err = createQrCode()
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

func (r *RealEstatePostgres) GetAll(userId int) (realEstates []model.RealEstate, err error) {
	rows, err := r.db.Query("SELECT id,user_profile_id, address, qr_code, created_at FROM real_estate WHERE user_profile_id = $1", userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var realEstate model.RealEstate
		err = rows.Scan(&realEstate.Id, &realEstate.UserProfileId, &realEstate.Address, &realEstate.QrCode, &realEstate.CreatedAt)
		if err != nil {
			return nil, err
		}
		realEstates = append(realEstates, realEstate)
	}
	return realEstates, nil
}

func (r *RealEstatePostgres) Get(userId int, id string) (realEstate model.RealEstate, err error) {
	err = r.db.QueryRow("SELECT id,user_profile_id, address, qr_code, created_at FROM real_estate WHERE user_profile_id = $1 AND id = $2", userId, id).Scan(&realEstate.Id, &realEstate.UserProfileId, &realEstate.Address, &realEstate.QrCode, &realEstate.CreatedAt)
	if err != nil {
		return model.RealEstate{}, err
	}
	return realEstate, nil
}

func (r *RealEstatePostgres) Delete(userId int, id string) error {
	var qrCode string
	err := r.db.QueryRow("SELECT qr_code FROM real_estate WHERE id = $1", id).Scan(&qrCode)
	if err != nil {
		return err
	}
	_, err = r.db.Exec("DELETE FROM real_estate WHERE id = $1 and user_profile_id = $2", id, userId)
	if err != nil {
		return err
	}

	err = deleteQRCode("./qr/" + qrCode + ".png")
	if err != nil {
		return err
	}

	return nil
}

func (r *RealEstatePostgres) Update(userId int, id string, realEstate *model.RealEstate) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	realEstate.QrCode, err = createQrCode()
	if err != nil {
		return err
	}

	var qrCode string
	err = tx.QueryRow("SELECT qr_code FROM real_estate WHERE id = $1", id).Scan(&qrCode)
	if err != nil {
		return err
	}

	err = deleteQRCode("./qr/" + qrCode + ".png")
	if err != nil {
		return err
	}

	_, err = tx.Exec("UPDATE real_estate SET address = $1, qr_code = $2 WHERE id = $3 AND user_profile_id = $4", realEstate.Address, realEstate.QrCode, id, userId)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (r *RealEstatePostgres) CheckAddress(address *string) error {
	return r.db.Select("SELECT * FROM real_estate WHERE address = $1", *address)
}

func deleteQRCode(pathFileName string) error {
	return os.Remove(pathFileName)
}

func createQrCode() (string, error) {
	qrCode := strconv.FormatInt(time.Now().Unix(), 10)
	data := "http://localhost:8080/api/visit/qr_code/" + qrCode

	err := writeQRCodeToFile("./qr/"+qrCode+".png", data)
	if err != nil {
		return "", err
	}
	return qrCode, nil
}

func writeQRCodeToFile(pathFileName, data string) error {
	return qrcode.WriteFile(data, qrcode.Medium, 256, pathFileName)
}
