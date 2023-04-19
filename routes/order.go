package routes

import (
	"e-commerce/handlers"
	"github.com/gorilla/mux"
	"sync"
)

func setupOrderRoutes(r mux.Router, wg *sync.WaitGroup) {
	r.HandleFunc("orders", handlers.ListOrders).Methods("GET")
	r.HandleFunc("order/{id}", handlers.GetOrder).Methods("GET")
	defer wg.Done()
}
