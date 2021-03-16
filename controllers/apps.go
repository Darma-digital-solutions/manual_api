package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/iamJune20/dds/models"
	_ "github.com/lib/pq"
)

type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

type Response struct {
	Status  int             `json:"status"`
	Message string          `json:"message"`
	Data    []models.OutNya `json:"data"`
}

func GetAllApp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	outNya, err := models.GetAllAppModel()

	if err != nil {
		log.Fatalf("Tidak bisa mengambil data. %v", err)
	}

	var response Response
	response.Status = 200
	response.Message = "Success"
	response.Data = outNya

	json.NewEncoder(w).Encode(response)
}
