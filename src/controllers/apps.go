package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/iamJune20/dds/config"
	"github.com/iamJune20/dds/src/modules/app/repository"

	"github.com/iamJune20/dds/src/modules/app/model"
	_ "github.com/lib/pq"
)

type response struct {
	Status  int    `json:"status"`
	Message string `json:"message,omitempty"`
	Code    string `json:"code,omitempty"`
}

type ResponseError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type ResponseGetApps struct {
	Status  int        `json:"status"`
	Message string     `json:"message"`
	Data    model.Apps `json:"data"`
}

type ResponseGetApp struct {
	Status  int        `json:"status"`
	Message string     `json:"message"`
	Data    *model.App `json:"data"`
}

func GetApps(w http.ResponseWriter, r *http.Request) {
	db := config.CreateConnection()
	appRepository := repository.NewAppRepository(db)

	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	out, err := appRepository.FindAll()

	if err != nil {
		log.Fatalf("Error: %v", err)
		return
	}

	var response ResponseGetApps
	response.Status = 200
	response.Message = "Success"
	response.Data = out

	json.NewEncoder(w).Encode(response)
}

func GetApp(w http.ResponseWriter, r *http.Request) {
	db := config.CreateConnection()
	appRepository := repository.NewAppRepository(db)

	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	params := mux.Vars(r)

	out, err := appRepository.FindByID(params["code"])

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		var response ResponseError
		response.Status = 404
		response.Message = "Code Not Found"
		json.NewEncoder(w).Encode(response)
	} else {
		var response ResponseGetApp
		response.Status = 200
		response.Message = "Success"
		response.Data = out
		json.NewEncoder(w).Encode(response)
	}
}

func InsertApp(w http.ResponseWriter, r *http.Request) {
	db := config.CreateConnection()
	appRepository := repository.NewAppRepository(db)

	var app *model.App
	app = model.NewApp()
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&app)

	if err != nil {
		fmt.Printf("Error : %v", err)
	}
	defer r.Body.Close()

	out, err := appRepository.Save(app)
	if err != nil {
		log.Fatalf("Error : %v", err)
	}
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		var response ResponseError
		response.Status = 404
		response.Message = "Code Not Found"
		json.NewEncoder(w).Encode(response)
	} else {
		res := response{
			Status:  201,
			Message: "Success",
			Code:    out,
		}
		json.NewEncoder(w).Encode(res)
	}
}
