package routes

import (
	"e-commerce/controller"
	"github.com/gorilla/mux"
)

func setupUserRoutes(r *mux.Router, c controller.Controller) {
	r.HandleFunc("/register", c.Register).Methods("POST")
	//r.HandleFunc("/login", handlers.LoginUser).Methods("POST")
	//r.HandleFunc("/logout", handlers.LogoutUser).Methods("POST")
	//r.HandleFunc("/reset-password", handlers.ResetUserPassword).Methods("POST")
}
