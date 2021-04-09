package controllers

import (
	"encoding/json"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/gorilla/mux"
	"github.com/iamJune20/dds/config"
	appRe "github.com/iamJune20/dds/src/modules/app/repository"
	"github.com/iamJune20/dds/src/modules/search/model"
	"github.com/iamJune20/dds/src/modules/search/repository"
)

type ResponseGetSearchAll struct {
	Status  int             `json:"status"`
	Message string          `json:"message"`
	Data    model.SearchAll `json:"data"`
}

type ResponseGetSearchOne struct {
	Status  int              `json:"status"`
	Message string           `json:"message"`
	Data    *model.SearchOne `json:"data"`
}

func GetSearchAll(w http.ResponseWriter, r *http.Request) {
	db := config.CreateConnection()
	searchRepository := repository.NewSearchRepository(db)

	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// out, err := searchRepository.FindAll()
	errSearch := validation.Validate(r.FormValue("Search"),
		validation.Required,
	)
	if errSearch != nil {
		w.WriteHeader(http.StatusNotFound)
		ReturnError(w, 404, "Field search harus diisi")
		return
	}
	out, err := searchRepository.FindAll(r.FormValue("Search"))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		var response ResponseError
		response.Status = 404
		response.Message = "Not Found"
		json.NewEncoder(w).Encode(response)
		return
	}

	var response ResponseGetSearchAll
	response.Status = 200
	response.Message = "Success"
	response.Data = out

	json.NewEncoder(w).Encode(response)
}

func GetSearchOne(w http.ResponseWriter, r *http.Request) {
	db := config.CreateConnection()
	searchRepository := repository.NewSearchRepository(db)

	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	params := mux.Vars(r)
	errSearch := validation.Validate(r.FormValue("Search"),
		validation.Required,
	)
	if errSearch != nil {
		w.WriteHeader(http.StatusNotFound)
		ReturnError(w, 404, "Field search harus diisi")
		return
	}
	errAppCode := validation.Validate(params["app_code"],
		validation.Required,
	)
	if errAppCode != nil {
		w.WriteHeader(http.StatusBadRequest)
		ReturnError(w, 400, "AppCode harus diisi")
		return
	}
	appRepository := appRe.NewAppRepository(db)
	app, err := appRepository.FindByID(params["app_code"])

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		var response ResponseError
		response.Status = 404
		response.Message = "Code Not Found"
		json.NewEncoder(w).Encode(response)
		return
	} else {
		if app == nil {
			w.WriteHeader(http.StatusNotFound)
			var response ResponseError
			response.Status = 404
			response.Message = "Code Not Found"
			json.NewEncoder(w).Encode(response)
			return
		}
	}
	out, err := searchRepository.FindOne(params["app_code"], r.FormValue("Search"))

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		var response ResponseError
		response.Status = 404
		response.Message = "Content Not Found"
		json.NewEncoder(w).Encode(response)
		return
	}

	var response ResponseGetSearchOne
	response.Status = 200
	response.Message = "Success"
	response.Data = out

	json.NewEncoder(w).Encode(response)
}
