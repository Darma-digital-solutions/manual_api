package models

import (
	"log"

	"github.com/iamJune20/dds/config"

	_ "github.com/lib/pq"
)

type OutNya struct {
	appCode  string `json:"appCode"`
	appName  string `json:"appName"`
	appLogo  string `json:"appLogo"`
	createAt string `json:"createAt"`
}

func getAllAppModel() ([]OutNya, error) {
	db := config.CreateConnection()

	defer db.Close()

	var outNya []OutNya

	sqlStatement := `SELECT * FROM apps`

	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("tidak bisa mengeksekusi query. %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var apps OutNya

		err = rows.Scan(&apps.appCode, &apps.appName, &apps.appLogo, &apps.createAt)

		if err != nil {
			log.Fatalf("tidak bisa mengambil data. %v", err)
		}

		outNya = append(outNya, apps)
	}

	// return empty buku atau jika error
	return outNya, err
}
