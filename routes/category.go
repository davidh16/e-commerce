package routes

import (
	"e-commerce/controller"
	"e-commerce/handlers"
	"github.com/gorilla/mux"
)

func setupCategoryRoutes(r *mux.Router, c controller.Controller) {
	r.HandleFunc("/category", handlers.CreateCategory).Methods("POST")
	r.HandleFunc("/category/update/{id}", handlers.UpdateCategory).Methods("POST")
	r.HandleFunc("/category/delete/{id}", handlers.DeleteCategory).Methods("DELETE")
	r.HandleFunc("/category/{id}", handlers.GetCategory).Methods("GET")
	r.HandleFunc("/categories", handlers.ListCategories).Methods("GET")
}
