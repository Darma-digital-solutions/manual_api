package models

import (
	"log"

	"github.com/iamJune20/dds/config"

	_ "github.com/lib/pq"
)

type OutNya struct {
	app_code  string `json:"app_code"`
	app_name  string `json:"app_name"`
	app_logo  string `json:"app_logo"`
	create_at string `json:"create_at"`
}

func GetAllAppModel() ([]OutNya, error) {
	db := config.CreateConnection()

	defer db.Close()

	var outNya []OutNya

	sqlStatement := `SELECT * FROM public.apps WHERE apps.app_publish = 'Yes'`

	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("tidak bisa mengeksekusi query. %v", err)
	}

	defer rows.Close()
	for rows.Next() {
		var apps OutNya

		err = rows.Scan(&apps.app_code, &apps.app_name, &apps.app_logo, &apps.create_at)

		if err != nil {
			log.Fatalf("tidak bisa mengambil data. %v", err)
		}
		outNya = append(outNya, apps)
	}

	// return empty buku atau jika error
	return outNya, err
}
