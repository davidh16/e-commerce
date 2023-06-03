package routes

import (
	"e-commerce/controller"
	"github.com/gorilla/mux"
)

func setupOrderRoutes(r *mux.Router, c controller.Controller) {
	r.HandleFunc("orders", c.ListOrders).Methods("GET")
	r.HandleFunc("order/{id}", c.GetOrder).Methods("GET")
}
