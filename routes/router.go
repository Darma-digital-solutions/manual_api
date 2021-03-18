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
	// router.HandleFunc("/api/buku/{id}", controller.UpdateBuku).Methods("PUT")
	// router.HandleFunc("/api/buku/{id}", controller.HapusBuku).Methods("DELETE")

	return router
}
