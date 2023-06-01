package routes

import (
	"e-commerce/controller"
	"e-commerce/handlers"
	"github.com/gorilla/mux"
)

func setupOrderRoutes(r *mux.Router, c controller.Controller) {
	r.HandleFunc("orders", handlers.ListOrders).Methods("GET")
	r.HandleFunc("order/{id}", handlers.GetOrder).Methods("GET")
}
