package routes

import (
	"github.com/gorilla/mux"
	"github.com/iamJune20/dds/controllers"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/apps", controllers.GetAllApp).Methods("GET", "OPTIONS")
	// router.HandleFunc("/api/buku/{id}", controller.AmbilBuku).Methods("GET", "OPTIONS")
	// router.HandleFunc("/api/buku", controller.TmbhBuku).Methods("POST", "OPTIONS")
	// router.HandleFunc("/api/buku/{id}", controller.UpdateBuku).Methods("PUT", "OPTIONS")
	// router.HandleFunc("/api/buku/{id}", controller.HapusBuku).Methods("DELETE", "OPTIONS")

	return router
}
