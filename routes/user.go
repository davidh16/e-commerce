package routes

import (
	"e-commerce/handlers"
	"github.com/gorilla/mux"
	"sync"
)

func setupUserRoutes(r *mux.Router, wg *sync.WaitGroup) {
	r.HandleFunc("/register", handlers.RegisterUser).Methods("POST")
	r.HandleFunc("/login", handlers.LoginUser).Methods("POST")
	r.HandleFunc("/logout", handlers.LogoutUser).Methods("POST")
	r.HandleFunc("/reset-password", handlers.ResetUserPassword).Methods("POST")
	defer wg.Done()
}
