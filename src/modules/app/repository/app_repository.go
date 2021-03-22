package repository

import (
	"database/sql"
	"time"

	"github.com/iamJune20/dds/src/modules/app/model"
)

type appRepository struct {
	db *sql.DB
}

func NewAppRepository(db *sql.DB) *appRepository {
	return &appRepository{db}
}

func (r *appRepository) FindAll() (model.Apps, error) {

	query := `SELECT * FROM "apps" WHERE "app_publish" = 'Yes' AND "delete_at" IS NULL`

	var apps model.Apps

	rows, err := r.db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var app model.App

		err = rows.Scan(&app.Code, &app.Name, &app.Logo, &app.CreatedAt, &app.UpdatedAt, &app.DeleteAt, &app.Publish)

		if err != nil {
			return nil, err
		}

		apps = append(apps, app)
	}

	return apps, nil
}

func (r *appRepository) FindByID(app_code string) (*model.App, error) {

	query := `SELECT * FROM "apps" WHERE "app_publish" = 'Yes' AND "delete_at" IS NULL AND "app_code" = $1`

	var app model.App

	statement, err := r.db.Prepare(query)

	if err != nil {
		return nil, err
	}

	defer statement.Close()

	err = statement.QueryRow(app_code).Scan(&app.Code, &app.Name, &app.Logo, &app.CreatedAt, &app.UpdatedAt, &app.DeleteAt, &app.Publish)

	if err != nil {
		return nil, err
	}

	return &app, nil
}

func (r *appRepository) Save(app *model.App) (string, error) {

	query := `INSERT INTO apps (app_name, app_logo,create_at,update_at, app_publish) VALUES ($1, $2, $3, $4, 'Yes') RETURNING app_code`

	var Code string
	err := r.db.QueryRow(query, app.Name, app.Logo, app.CreatedAt, app.UpdatedAt).Scan(&Code)

	if err != nil {
		return "Data gagal disimpan", err
	}

	return Code, err
}

func (r *appRepository) Update(app_code string, app *model.App) (string, error) {
	query := `UPDATE apps SET app_name=$1, app_logo=$2, update_at=$3 WHERE app_code=$4 AND "app_publish" = 'Yes' AND "delete_at" IS NULL`

	statement, err := r.db.Prepare(query)

	if err != nil {
		return "Data gagal di ubah", err
	}

	defer statement.Close()

	_, err = statement.Exec(app.Name, app.Logo, app.UpdatedAt, app_code)

	if err != nil {
		return "Data gagal di ubah", err
	}

	return "Data berhasil diubah", err
}

func (r *appRepository) Delete(app_code string) (string, error) {
	query := `UPDATE apps SET delete_at=$1, app_publish='No' WHERE app_code=$2 AND "app_publish" = 'Yes' AND "delete_at" IS NULL`

	statement, err := r.db.Prepare(query)

	if err != nil {
		return "Data gagal di hapus", err
	}

	defer statement.Close()

	_, err = statement.Exec(time.Now(), app_code)

	if err != nil {
		return "Data gagal di hapus", err
	}

	return "Data berhasil dihapus", err
}
