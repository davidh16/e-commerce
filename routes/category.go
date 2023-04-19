package routes

import (
	"e-commerce/handlers"
	"github.com/gorilla/mux"
	"sync"
)

func setupCategoryRoutes(r mux.Router, wg *sync.WaitGroup) {
	r.HandleFunc("/category", handlers.CreateCategory).Methods("POST")
	r.HandleFunc("/category/update/{id}", handlers.UpdateCategory).Methods("POST")
	r.HandleFunc("/category/delete/{id}", handlers.DeleteCategory).Methods("DELETE")
	r.HandleFunc("/category/{id}", handlers.GetCategory).Methods("GET")
	r.HandleFunc("/categories", handlers.ListCategories).Methods("GET")
	defer wg.Done()
}
