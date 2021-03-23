package routes

import (
	"github.com/gorilla/mux"
	"github.com/iamJune20/dds/src/controllers"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/apps", controllers.GetApps).Methods("GET")
	router.HandleFunc("/api/app/{code}", controllers.GetApp).Methods("GET")
	router.HandleFunc("/api/app", controllers.InsertApp).Methods("POST")
	router.HandleFunc("/api/app/{code}", controllers.UpdateApp).Methods("PUT")
	router.HandleFunc("/api/app/{code}", controllers.DeleteApp).Methods("DELETE")

	router.HandleFunc("/api/categories", controllers.GetCategories).Methods("GET")
	router.HandleFunc("/api/categoriesByApp/{code}", controllers.GetCategoriesByAppCode).Methods("GET")
	router.HandleFunc("/api/category/{code}", controllers.GetCategory).Methods("GET")
	router.HandleFunc("/api/category/{app_code}", controllers.InsertCategory).Methods("POST")
	router.HandleFunc("/api/category/{code}", controllers.UpdateCategory).Methods("PUT")
	router.HandleFunc("/api/category/{code}", controllers.DeleteCategory).Methods("DELETE")

	router.HandleFunc("/api/manuals", controllers.GetManuals).Methods("GET")
	router.HandleFunc("/api/manualsByApp/{code}", controllers.GetManualsByAppCode).Methods("GET")
	router.HandleFunc("/api/manual/{code}", controllers.GetManual).Methods("GET")
	router.HandleFunc("/api/manual/{app_code}", controllers.InsertManual).Methods("POST")
	router.HandleFunc("/api/manual/{code}", controllers.UpdateManual).Methods("PUT")
	router.HandleFunc("/api/manual/{code}", controllers.DeleteManual).Methods("DELETE")

	router.HandleFunc("/api/contents", controllers.GetContents).Methods("GET")
	router.HandleFunc("/api/content/{code}", controllers.GetContent).Methods("GET")
	router.HandleFunc("/api/content/{manual_code}/{category_code}", controllers.InsertContent).Methods("POST")
	router.HandleFunc("/api/content/{code}", controllers.UpdateContent).Methods("PUT")
	router.HandleFunc("/api/content/{code}", controllers.DeleteContent).Methods("DELETE")
	router.HandleFunc("/api/contentsIn/{manual_code}/{category_code}", controllers.GetContentInManualAndCategory).Methods("GET")

	router.HandleFunc("/api/searchAll", controllers.GetSearchAll).Methods("POST")
	router.HandleFunc("/api/searchOne/{app_code}", controllers.GetSearchOne).Methods("POST")

	// GetSearchAll
	return router
}
