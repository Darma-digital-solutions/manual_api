package repository

import (
	"database/sql"
	"log"

	"github.com/iamJune20/dds/src/modules/app/model"
)

type appRepository struct {
	db *sql.DB
}

func NewAppRepository(db *sql.DB) *appRepository {
	return &appRepository{db}
}

func (r *appRepository) FindAll() (model.Apps, error) {

	query := `SELECT * FROM "apps" WHERE "app_publish" = 'Yes'`

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

	query := `SELECT * FROM "apps" WHERE "app_publish" = 'Yes' AND "app_code" = $1`

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
		log.Fatalf("Tidak Bisa mengeksekusi query. %v", err)
	}

	return Code, err
}
