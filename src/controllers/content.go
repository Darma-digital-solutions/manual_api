package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/gorilla/mux"
	"github.com/iamJune20/dds/config"
	categoryRe "github.com/iamJune20/dds/src/modules/category/repository"
	"github.com/iamJune20/dds/src/modules/content/repository"
	manualRe "github.com/iamJune20/dds/src/modules/manual/repository"

	"github.com/iamJune20/dds/src/modules/content/model"
	_ "github.com/lib/pq"
)

type ResponseGetContents struct {
	Status  int            `json:"status"`
	Message string         `json:"message"`
	Data    model.Contents `json:"data"`
}

type ResponseGetContent struct {
	Status  int            `json:"status"`
	Message string         `json:"message"`
	Data    *model.Content `json:"data"`
}

func GetContents(w http.ResponseWriter, r *http.Request) {
	db := config.CreateConnection()
	contentRepository := repository.NewContentRepository(db)

	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	out, err := contentRepository.FindAll()

	if err != nil {
		log.Fatalf("Error: %v", err)
		return
	}

	var response ResponseGetContents
	response.Status = 200
	response.Message = "Success"
	response.Data = out

	json.NewEncoder(w).Encode(response)
}

func GetContent(w http.ResponseWriter, r *http.Request) {
	db := config.CreateConnection()
	contentRepository := repository.NewContentRepository(db)

	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	params := mux.Vars(r)

	out, err := contentRepository.FindByID(params["code"])

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		var response ResponseError
		response.Status = 404
		response.Message = "Code Not Found"
		json.NewEncoder(w).Encode(response)
	} else {
		var response ResponseGetContent
		response.Status = 200
		response.Message = "Success"
		response.Data = out
		json.NewEncoder(w).Encode(response)
	}
}

func GetContentInManualAndCategory(w http.ResponseWriter, r *http.Request) {
	db := config.CreateConnection()
	contentRepository := repository.NewContentRepository(db)

	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	params := mux.Vars(r)
	manualRepository := manualRe.NewManualRepository(db)
	manual, err := manualRepository.FindByID(params["manual_code"])

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		var response ResponseError
		response.Status = 404
		response.Message = "Manual code Not Found"
		json.NewEncoder(w).Encode(response)
		return
	} else {
		if manual == nil {
			w.WriteHeader(http.StatusNotFound)
			var response ResponseError
			response.Status = 404
			response.Message = "Manual code Not Found"
			json.NewEncoder(w).Encode(response)
			return
		}
	}

	categoryRepository := categoryRe.NewCategoryRepository(db)
	category, err := categoryRepository.FindByID(params["category_code"])

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		var response ResponseError
		response.Status = 404
		response.Message = "Category code Not Found"
		json.NewEncoder(w).Encode(response)
		return
	} else {
		if category == nil {
			w.WriteHeader(http.StatusNotFound)
			var response ResponseError
			response.Status = 404
			response.Message = "Category code Not Found"
			json.NewEncoder(w).Encode(response)
			return
		}
	}
	out, err := contentRepository.FindOne(params["manual_code"], params["category_code"])

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		var response ResponseError
		response.Status = 404
		response.Message = "Code Not Found"
		json.NewEncoder(w).Encode(response)
	} else {
		var response ResponseGetContents
		response.Status = 200
		response.Message = "Success"
		response.Data = out
		json.NewEncoder(w).Encode(response)
	}
}

func InsertContent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	db := config.CreateConnection()
	contentRepository := repository.NewContentRepository(db)

	var content *model.Content
	content = model.NewContent()
	content.Title = r.FormValue("Title")
	errTitle := validation.Validate(content.Title,
		validation.Required,
	)
	if errTitle != nil {
		w.WriteHeader(http.StatusBadRequest)
		ReturnError(w, 400, "Judul konten harus diisi")
		return
	}
	content.Desc = r.FormValue("Desc")
	errDesc := validation.Validate(content.Desc,
		validation.Required,
	)
	if errDesc != nil {
		w.WriteHeader(http.StatusBadRequest)
		ReturnError(w, 400, "Deskripsi konten harus diisi")
		return
	}
	params := mux.Vars(r)
	manualRepository := manualRe.NewManualRepository(db)
	manual, err := manualRepository.FindByID(params["manual_code"])

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		var response ResponseError
		response.Status = 404
		response.Message = "Manual code Not Found"
		json.NewEncoder(w).Encode(response)
		return
	} else {
		if manual == nil {
			w.WriteHeader(http.StatusNotFound)
			var response ResponseError
			response.Status = 404
			response.Message = "Manual code Not Found"
			json.NewEncoder(w).Encode(response)
			return
		} else {
			content.ManualCode = params["manual_code"]
		}
	}

	categoryRepository := categoryRe.NewCategoryRepository(db)
	category, err := categoryRepository.FindByID(params["category_code"])

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		var response ResponseError
		response.Status = 404
		response.Message = "Category code Not Found"
		json.NewEncoder(w).Encode(response)
		return
	} else {
		if category == nil {
			w.WriteHeader(http.StatusNotFound)
			var response ResponseError
			response.Status = 404
			response.Message = "Category code Not Found"
			json.NewEncoder(w).Encode(response)
			return
		} else {
			content.CategoryCode = params["category_code"]
		}
	}
	out, err := contentRepository.Save(content)
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

func UpdateContent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	db := config.CreateConnection()
	contentRepository := repository.NewContentRepository(db)

	var content *model.Content
	content = model.NewContent()
	content.Title = r.FormValue("Title")
	errTitle := validation.Validate(content.Title,
		validation.Required,
	)
	if errTitle != nil {
		w.WriteHeader(http.StatusBadRequest)
		ReturnError(w, 400, "Judul konten harus diisi")
		return
	}
	content.Desc = r.FormValue("Desc")
	errDesc := validation.Validate(content.Desc,
		validation.Required,
	)
	if errDesc != nil {
		w.WriteHeader(http.StatusBadRequest)
		ReturnError(w, 400, "Deskripsi konten harus diisi")
		return
	}
	params := mux.Vars(r)
	errCode := validation.Validate(params["code"],
		validation.Required,
	)
	if errCode != nil {
		w.WriteHeader(http.StatusNotFound)
		ReturnError(w, 404, "Kode tidak ditemukan")
		return
	}
	manualRepository := manualRe.NewManualRepository(db)
	manual, err := manualRepository.FindByID(r.FormValue("ManualCode"))

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		var response ResponseError
		response.Status = 404
		response.Message = "Manual code Not Found"
		json.NewEncoder(w).Encode(response)
		return
	} else {
		if manual == nil {
			w.WriteHeader(http.StatusNotFound)
			var response ResponseError
			response.Status = 404
			response.Message = "Manual code Not Found"
			json.NewEncoder(w).Encode(response)
			return
		} else {
			content.ManualCode = r.FormValue("ManualCode")
		}
	}

	categoryRepository := categoryRe.NewCategoryRepository(db)
	category, err := categoryRepository.FindByID(r.FormValue("CategoryCode"))

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		var response ResponseError
		response.Status = 404
		response.Message = "Category code Not Found"
		json.NewEncoder(w).Encode(response)
		return
	} else {
		if category == nil {
			w.WriteHeader(http.StatusNotFound)
			var response ResponseError
			response.Status = 404
			response.Message = "Category code Not Found"
			json.NewEncoder(w).Encode(response)
			return
		} else {
			content.CategoryCode = r.FormValue("CategoryCode")
		}
	}
	out, err := contentRepository.Update(params["code"], content)
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

func DeleteContent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	db := config.CreateConnection()
	contentRepository := repository.NewContentRepository(db)

	params := mux.Vars(r)
	errCode := validation.Validate(params["code"],
		validation.Required,
	)
	if errCode != nil {
		w.WriteHeader(http.StatusNotFound)
		ReturnError(w, 404, "Kode tidak ditemukan")
		return
	}
	out, err := contentRepository.Delete(params["code"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		ReturnError(w, 400, "Gagal menghapus data")
		return
	} else {
		res := response{
			Status:  200,
			Message: out,
			Code:    params["code"],
		}
		json.NewEncoder(w).Encode(res)
	}
}
