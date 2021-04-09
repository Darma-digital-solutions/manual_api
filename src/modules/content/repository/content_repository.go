package repository

import (
	"database/sql"
	"time"

	"github.com/iamJune20/dds/src/modules/content/model"
)

type contentRepository struct {
	db *sql.DB
}

func NewContentRepository(db *sql.DB) *contentRepository {
	return &contentRepository{db}
}

func (r *contentRepository) FindAll() (model.Contents, error) {

	query := `SELECT * FROM "content" WHERE "content_publish" = 'Yes' AND "delete_at" IS NULL`

	var contents model.Contents

	rows, err := r.db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var content model.Content

		err = rows.Scan(&content.Code, &content.Title, &content.Desc, &content.ManualCode, &content.CategoryCode, &content.CreatedAt, &content.UpdatedAt, &content.DeleteAt, &content.Publish)

		if err != nil {
			return nil, err
		}

		contents = append(contents, content)
	}

	return contents, nil
}
func (r *contentRepository) FindOne(manual_code string, category_code string) (model.Contents, error) {

	query := `SELECT * FROM "content" WHERE "content_publish" = 'Yes' AND "delete_at" IS NULL AND "manual_code" = $1 AND "category_code" = $2`

	var contents model.Contents

	rows, err := r.db.Query(query, manual_code, category_code)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var content model.Content

		err = rows.Scan(&content.Code, &content.Title, &content.Desc, &content.CategoryCode, &content.ManualCode, &content.CreatedAt, &content.UpdatedAt, &content.DeleteAt, &content.Publish)

		if err != nil {
			return nil, err
		}

		contents = append(contents, content)
	}

	return contents, nil
}
func (r *contentRepository) FindByID(content_code string) (*model.Content, error) {

	query := `SELECT * FROM "content" WHERE "content_publish" = 'Yes' AND "delete_at" IS NULL AND "content_code" = $1`

	var content model.Content

	statement, err := r.db.Prepare(query)

	if err != nil {
		return nil, err
	}

	defer statement.Close()

	err = statement.QueryRow(content_code).Scan(&content.Code, &content.Title, &content.Desc, &content.ManualCode, &content.CategoryCode, &content.CreatedAt, &content.UpdatedAt, &content.DeleteAt, &content.Publish)

	if err != nil {
		return nil, err
	}

	return &content, nil
}

func (r *contentRepository) Save(content *model.Content) (string, error) {

	query := `INSERT INTO content (content_title, content_desc,category_code,manual_code,create_at,update_at, content_publish) VALUES ($1, $2, $3, $4, $5, $6,'Yes') RETURNING content_code`

	var Code string
	err := r.db.QueryRow(query, content.Title, content.Desc, content.CategoryCode, content.ManualCode, content.CreatedAt, content.UpdatedAt).Scan(&Code)

	if err != nil {
		return "Data gagal disimpan", err
	}

	return Code, err
}

func (r *contentRepository) Update(content_code string, content *model.Content) (string, error) {
	query := `UPDATE content SET content_title=$1, content_desc=$2, category_code=$3, manual_code=$4, update_at=$5 WHERE content_code=$6 AND "content_publish" = 'Yes' AND "delete_at" IS NULL`

	statement, err := r.db.Prepare(query)

	if err != nil {
		// fmt.Printf("Error : %v", err)
		return "Data gagal di ubah", err
	}

	defer statement.Close()

	_, err = statement.Exec(content.Title, content.Desc, content.CategoryCode, content.ManualCode, content.UpdatedAt, content_code)

	if err != nil {
		return "Data gagal di ubah", err
	}

	return "Data berhasil diubah", err
}

func (r *contentRepository) Delete(content_code string) (string, error) {
	query := `UPDATE content SET delete_at=$1, content_publish='No' WHERE content_code=$2 AND "content_publish" = 'Yes' AND "delete_at" IS NULL`

	statement, err := r.db.Prepare(query)

	if err != nil {
		return "Data gagal di hapus", err
	}

	defer statement.Close()

	_, err = statement.Exec(time.Now(), content_code)

	if err != nil {
		return "Data gagal di hapus", err
	}

	return "Data berhasil dihapus", err
}
