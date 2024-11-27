package controllers

import (
	"encoding/json"
	"fmt"
	"go_learn/mongoDB/pkg/models"
	"go_learn/mongoDB/pkg/utils"
	"net/http"

	"github.com/gorilla/mux"
)

func GetCars(w http.ResponseWriter, r *http.Request) {
	cars := models.GetAllCars()
	res, _ := json.Marshal(cars)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetCarById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	carId := vars["carId"]
	carDetails, err := models.GetCarById(carId)

	if err != nil {
		fmt.Println("Error while fetching car by ID:", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	res, _ := json.Marshal(carDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateCar(w http.ResponseWriter, r *http.Request) {
	var car models.Car
	utils.ParseBody(r, &car)

	car, err := models.CreateCar(car)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	res, _ := json.Marshal(car)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateCar(w http.ResponseWriter, r *http.Request) {
	var updateCar models.Car
	utils.ParseBody(r, &updateCar)
	vars := mux.Vars(r)
	carId := vars["carId"]

	car, err := models.UpdateCar(carId, updateCar)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	res, _ := json.Marshal(car)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteCar(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	carId := vars["carId"]

	err := models.DeleteCar(carId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}
