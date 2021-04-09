package repository

import (
	"database/sql"
	"time"

	"github.com/iamJune20/dds/src/modules/manual/model"
)

type manualRepository struct {
	db *sql.DB
}

func NewManualRepository(db *sql.DB) *manualRepository {
	return &manualRepository{db}
}

func (r *manualRepository) FindAll() (model.Manuals, error) {

	query := `SELECT * FROM "manual" WHERE "manual_publish" = 'Yes' AND "delete_at" IS NULL`

	var manuals model.Manuals

	rows, err := r.db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var manual model.Manual

		err = rows.Scan(&manual.Code, &manual.Name, &manual.Desc, &manual.Icon, &manual.AppCode, &manual.CreatedAt, &manual.UpdatedAt, &manual.DeleteAt, &manual.Publish)

		if err != nil {
			return nil, err
		}

		manuals = append(manuals, manual)
	}

	return manuals, nil
}
func (r *manualRepository) FindByAppCode(app_code string) (model.Manuals, error) {

	query := `SELECT * FROM "manual" WHERE "manual_publish" = 'Yes' AND "delete_at" IS NULL AND "app_code" = $1`

	var manuals model.Manuals

	rows, err := r.db.Query(query, app_code)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var manual model.Manual

		err = rows.Scan(&manual.Code, &manual.Name, &manual.Desc, &manual.Icon, &manual.AppCode, &manual.CreatedAt, &manual.UpdatedAt, &manual.DeleteAt, &manual.Publish)

		if err != nil {
			return nil, err
		}

		manuals = append(manuals, manual)
	}

	return manuals, nil
}
func (r *manualRepository) FindByID(manual_code string) (*model.Manual, error) {

	query := `SELECT * FROM "manual" WHERE "manual_publish" = 'Yes' AND "delete_at" IS NULL AND "manual_code" = $1`

	var manual model.Manual

	statement, err := r.db.Prepare(query)

	if err != nil {
		return nil, err
	}

	defer statement.Close()

	err = statement.QueryRow(manual_code).Scan(&manual.Code, &manual.Name, &manual.Desc, &manual.Icon, &manual.AppCode, &manual.CreatedAt, &manual.UpdatedAt, &manual.DeleteAt, &manual.Publish)

	if err != nil {
		return nil, err
	}

	return &manual, nil
}

func (r *manualRepository) Save(manual *model.Manual) (string, error) {

	query := `INSERT INTO manual (manual_name, manual_desc,manual_icon,app_code,create_at,update_at, manual_publish) VALUES ($1, $2, $3, $4, $5, $6,'Yes') RETURNING manual_code`

	var Code string
	err := r.db.QueryRow(query, manual.Name, manual.Desc, manual.Icon, manual.AppCode, manual.CreatedAt, manual.UpdatedAt).Scan(&Code)

	if err != nil {
		return "Data gagal disimpan", err
	}

	return Code, err
}

func (r *manualRepository) Update(manual_code string, manual *model.Manual) (string, error) {
	query := `UPDATE manual SET manual_name=$1, manual_desc=$2, manual_icon=$3, app_code=$4, update_at=$5 WHERE manual_code=$6 AND "manual_publish" = 'Yes' AND "delete_at" IS NULL`

	statement, err := r.db.Prepare(query)

	if err != nil {
		return "Data gagal di ubah", err
	}

	defer statement.Close()

	_, err = statement.Exec(manual.Name, manual.Desc, manual.Icon, manual.AppCode, manual.UpdatedAt, manual_code)

	if err != nil {
		return "Data gagal di ubah", err
	}

	return "Data berhasil diubah", err
}

func (r *manualRepository) Delete(manual_code string) (string, error) {
	query := `UPDATE manual SET delete_at=$1, manual_publish='No' WHERE manual_code=$2 AND "manual_publish" = 'Yes' AND "delete_at" IS NULL`

	statement, err := r.db.Prepare(query)

	if err != nil {
		return "Data gagal di hapus", err
	}

	defer statement.Close()

	_, err = statement.Exec(time.Now(), manual_code)

	if err != nil {
		return "Data gagal di hapus", err
	}

	return "Data berhasil dihapus", err
}
