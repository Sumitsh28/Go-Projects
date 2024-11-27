package routes

import (
	"go_learn/mongoDB/pkg/controllers"

	"github.com/gorilla/mux"
)

var RegisterCarRoutes = func(router *mux.Router) {
	router.HandleFunc("/cars/", controllers.CreateCar).Methods("POST")
	router.HandleFunc("/cars/", controllers.GetCars).Methods("GET")
	router.HandleFunc("/cars/{carId}", controllers.GetCarById).Methods("GET")
	router.HandleFunc("/cars/{carId}", controllers.UpdateCar).Methods("PUT")
	router.HandleFunc("/cars/{carId}", controllers.DeleteCar).Methods("DELETE")
}
