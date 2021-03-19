package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation"

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

func ReturnError(w http.ResponseWriter, code int, message string) {
	var response ResponseError
	response.Status = code
	response.Message = message
	json.NewEncoder(w).Encode(response)
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
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	db := config.CreateConnection()
	appRepository := repository.NewAppRepository(db)

	var app *model.App
	app = model.NewApp()
	app.Name = r.FormValue("Name")
	errName := validation.Validate(app.Name,
		validation.Required,
	)
	if errName != nil {
		w.WriteHeader(http.StatusBadRequest)
		ReturnError(w, 400, "Nama aplikasi harus diisi")
		return
	}
	app.Logo = r.FormValue("Logo")
	errLogo := validation.Validate(app.Logo,
		validation.Required,
	)
	if errLogo != nil {
		w.WriteHeader(http.StatusBadRequest)
		ReturnError(w, 400, "Logo aplikasi harus diisi")
		return
	}
	out, err := appRepository.Save(app)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		ReturnError(w, 400, out)
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		ReturnError(w, 404, "Kode tidak ditemukan")
		return
	} else {
		res := response{
			Status:  201,
			Message: "Success",
			Code:    out,
		}
		json.NewEncoder(w).Encode(res)
	}
}

func UpdateApp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	db := config.CreateConnection()
	appRepository := repository.NewAppRepository(db)

	params := mux.Vars(r)
	errCode := validation.Validate(params["code"],
		validation.Required,
	)
	if errCode != nil {
		w.WriteHeader(http.StatusNotFound)
		ReturnError(w, 404, "Kode tidak ditemukan")
	}
	var app *model.App
	app = model.UpdateApp()
	app.Name = r.FormValue("Name")
	errName := validation.Validate(app.Name,
		validation.Required,
	)
	if errName != nil {
		w.WriteHeader(http.StatusBadRequest)
		ReturnError(w, 400, "Nama aplikasi harus diisi")
	}
	app.Logo = r.FormValue("Logo")
	errLogo := validation.Validate(app.Logo,
		validation.Required,
	)
	if errLogo != nil {
		w.WriteHeader(http.StatusBadRequest)
		ReturnError(w, 400, "Logo aplikasi harus diisi")
	}
	out, err := appRepository.Update(params["code"], app)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		ReturnError(w, 400, "Gagal mengubah data")
	} else {
		res := response{
			Status:  200,
			Message: out,
			Code:    params["code"],
		}
		json.NewEncoder(w).Encode(res)
	}
}

func DeleteApp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	db := config.CreateConnection()
	appRepository := repository.NewAppRepository(db)

	params := mux.Vars(r)
	errCode := validation.Validate(params["code"],
		validation.Required,
	)
	if errCode != nil {
		w.WriteHeader(http.StatusNotFound)
		ReturnError(w, 404, "Kode tidak ditemukan")
	}
	out, err := appRepository.Delete(params["code"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		ReturnError(w, 400, "Gagal menghapus data")
	} else {
		res := response{
			Status:  200,
			Message: out,
			Code:    params["code"],
		}
		json.NewEncoder(w).Encode(res)
	}
}
