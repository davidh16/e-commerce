package routes

import (
	"e-commerce/handlers"
	"github.com/gorilla/mux"
	"sync"
)

func setupProductRoutes(r mux.Router, wg *sync.WaitGroup) {
	r.HandleFunc("/products", handlers.ListProducts).Methods("GET")
	r.HandleFunc("/product/{id}", handlers.GetProduct).Methods("GET")
	r.HandleFunc("/product", handlers.CreateProduct).Methods("POST")
	r.HandleFunc("/product/update/{id}", handlers.UpdateProduct).Methods("POST")
	r.HandleFunc("/product/delete/{id}", handlers.DeleteProduct).Methods("DELETE")
	defer wg.Done()
}
