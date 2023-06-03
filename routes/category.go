package routes

import (
	"e-commerce/controller"
	"github.com/gorilla/mux"
)

func setupCategoryRoutes(r *mux.Router, c controller.Controller) {
	r.HandleFunc("/category", c.CreateCategory).Methods("POST")
	r.HandleFunc("/category/update/{id}", c.UpdateCategory).Methods("POST")
	r.HandleFunc("/category/delete/{id}", c.DeleteCategory).Methods("DELETE")
	r.HandleFunc("/category/{id}", c.GetCategory).Methods("GET")
	r.HandleFunc("/categories", c.ListCategories).Methods("GET")
}
