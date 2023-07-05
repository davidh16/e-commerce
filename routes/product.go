package routes

import (
	"e-commerce/server"
	"github.com/gorilla/mux"
)

func setupProductRoutes(r *mux.Router, c server.Controller) {
	r.HandleFunc("/products", c.ListProducts).Methods("GET")
	r.HandleFunc("/product/{id}", c.GetProduct).Methods("GET")
	r.HandleFunc("/product", c.CreateProduct).Methods("POST")
	r.HandleFunc("/product/update/{id}", c.UpdateProduct).Methods("POST")
	r.HandleFunc("/product/delete/{id}", c.DeleteProduct).Methods("DELETE")
}
