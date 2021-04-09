package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/gorilla/mux"
	"github.com/iamJune20/dds/config"
	appRe "github.com/iamJune20/dds/src/modules/app/repository"
	"github.com/iamJune20/dds/src/modules/manual/repository"

	"github.com/iamJune20/dds/src/modules/manual/model"
	_ "github.com/lib/pq"
)

type ResponseGetManuals struct {
	Status  int           `json:"status"`
	Message string        `json:"message"`
	Data    model.Manuals `json:"data"`
}

type ResponseGetManual struct {
	Status  int           `json:"status"`
	Message string        `json:"message"`
	Data    *model.Manual `json:"data"`
}

func GetManuals(w http.ResponseWriter, r *http.Request) {
	db := config.CreateConnection()
	manualRepository := repository.NewManualRepository(db)

	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	out, err := manualRepository.FindAll()

	if err != nil {
		log.Fatalf("Error: %v", err)
		return
	}

	var response ResponseGetManuals
	response.Status = 200
	response.Message = "Success"
	response.Data = out

	json.NewEncoder(w).Encode(response)
}

func GetManual(w http.ResponseWriter, r *http.Request) {
	db := config.CreateConnection()
	manualRepository := repository.NewManualRepository(db)

	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	params := mux.Vars(r)

	out, err := manualRepository.FindByID(params["code"])

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		var response ResponseError
		response.Status = 404
		response.Message = "Code Not Found"
		json.NewEncoder(w).Encode(response)
	} else {
		var response ResponseGetManual
		response.Status = 200
		response.Message = "Success"
		response.Data = out
		json.NewEncoder(w).Encode(response)
	}
}

func GetManualsByAppCode(w http.ResponseWriter, r *http.Request) {
	db := config.CreateConnection()
	manualRepository := repository.NewManualRepository(db)

	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	params := mux.Vars(r)

	out, err := manualRepository.FindByAppCode(params["code"])

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		var response ResponseError
		response.Status = 404
		response.Message = "Code Not Found"
		json.NewEncoder(w).Encode(response)
	} else {
		var response ResponseGetManuals
		response.Status = 200
		response.Message = "Success"
		response.Data = out
		json.NewEncoder(w).Encode(response)
	}
}

func InsertManual(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	db := config.CreateConnection()
	manualRepository := repository.NewManualRepository(db)

	var manual *model.Manual
	manual = model.NewManual()
	manual.Name = r.FormValue("Name")
	errName := validation.Validate(manual.Name,
		validation.Required,
	)
	if errName != nil {
		w.WriteHeader(http.StatusBadRequest)
		ReturnError(w, 400, "Nama kategory harus diisi")
		return
	}
	manual.Desc = r.FormValue("Desc")
	errDesc := validation.Validate(manual.Desc,
		validation.Required,
	)
	if errDesc != nil {
		w.WriteHeader(http.StatusBadRequest)
		ReturnError(w, 400, "Deskripsi kategory harus diisi")
		return
	}

	manual.Icon = r.FormValue("Icon")
	errIcon := validation.Validate(manual.Icon,
		validation.Required,
	)
	if errIcon != nil {
		w.WriteHeader(http.StatusBadRequest)
		ReturnError(w, 400, "Icon kategory harus diisi")
		return
	}
	params := mux.Vars(r)
	appRepository := appRe.NewAppRepository(db)
	app, err := appRepository.FindByID(params["app_code"])
	// fmt.Printf("Error : %v", err)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		var response ResponseError
		response.Status = 404
		response.Message = "Code Not Found"
		json.NewEncoder(w).Encode(response)
	} else {
		if app == nil {
			w.WriteHeader(http.StatusNotFound)
			var response ResponseError
			response.Status = 404
			response.Message = "Code Not Found"
			json.NewEncoder(w).Encode(response)
		} else {
			manual.AppCode = params["app_code"]
		}
	}

	out, err := manualRepository.Save(manual)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		ReturnError(w, 400, out)
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

func UpdateManual(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	db := config.CreateConnection()
	manualRepository := repository.NewManualRepository(db)

	params := mux.Vars(r)

	var manual *model.Manual
	manual = model.NewManual()
	manual.Name = r.FormValue("Name")
	errName := validation.Validate(manual.Name,
		validation.Required,
	)
	if errName != nil {
		w.WriteHeader(http.StatusBadRequest)
		ReturnError(w, 400, "Nama kategory harus diisi")
		return
	}
	manual.Desc = r.FormValue("Desc")
	errDesc := validation.Validate(manual.Desc,
		validation.Required,
	)
	if errDesc != nil {
		w.WriteHeader(http.StatusBadRequest)
		ReturnError(w, 400, "Deskripsi kategory harus diisi")
		return
	}

	manual.Icon = r.FormValue("Icon")
	errIcon := validation.Validate(manual.Icon,
		validation.Required,
	)
	if errIcon != nil {
		w.WriteHeader(http.StatusBadRequest)
		ReturnError(w, 400, "Icon kategory harus diisi")
		return
	}

	manual.AppCode = r.FormValue("AppCode")
	errAppCode := validation.Validate(manual.AppCode,
		validation.Required,
	)
	if errAppCode != nil {
		w.WriteHeader(http.StatusBadRequest)
		ReturnError(w, 400, "AppCode harus diisi")
		return
	}
	appRepository := appRe.NewAppRepository(db)
	app, err := appRepository.FindByID(manual.AppCode)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		var response ResponseError
		response.Status = 404
		response.Message = "Code Not Found"
		json.NewEncoder(w).Encode(response)
	} else {
		if app == nil {
			w.WriteHeader(http.StatusNotFound)
			var response ResponseError
			response.Status = 404
			response.Message = "Code Not Found"
			json.NewEncoder(w).Encode(response)
		}
	}

	out, err := manualRepository.Update(params["code"], manual)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		ReturnError(w, 400, out)
		return
	} else {
		res := response{
			Status:  200,
			Message: "Success",
			Code:    out,
		}
		json.NewEncoder(w).Encode(res)
	}
}

func DeleteManual(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	db := config.CreateConnection()
	manualRepository := repository.NewManualRepository(db)

	params := mux.Vars(r)
	errCode := validation.Validate(params["code"],
		validation.Required,
	)
	if errCode != nil {
		w.WriteHeader(http.StatusNotFound)
		ReturnError(w, 404, "Kode tidak ditemukan")
	}
	out, err := manualRepository.Delete(params["code"])
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
