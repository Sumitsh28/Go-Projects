package router

import (
	middleware "go_learn/postgresql/pkg/middlewares"

	"github.com/gorilla/mux"
)

func PostRouter() *mux.Router{

	router:= mux.NewRouter()
	router.HandleFunc("/api/stock",middleware.GetAllStocks).Methods("GET","OPTIONS")
	router.HandleFunc("/api/stock/{id}",middleware.GetStock).Methods("GET","OPTIONS")
	router.HandleFunc("/api/newstock",middleware.CreateStock).Methods("POST","OPTIONS")
	router.HandleFunc("/api/stock/{id}",middleware.UpdateStock).Methods("PUT","OPTIONS")
	router.HandleFunc("/api/deletestock/{id}",middleware.DeleteStock).Methods("DELETE","OPTIONS")

	return router
}