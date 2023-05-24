package routes

import (
	"e-commerce/handlers"
	"github.com/gorilla/mux"
	"sync"
)

func setupUserRoutes(r *mux.Router, wg *sync.WaitGroup, userHandler *handlers.UserHandler) {
	r.Handle("/register", userHandler).Methods("POST")
	//r.HandleFunc("/login", handlers.LoginUser).Methods("POST")
	//r.HandleFunc("/logout", handlers.LogoutUser).Methods("POST")
	//r.HandleFunc("/reset-password", handlers.ResetUserPassword).Methods("POST")
	defer wg.Done()
}
