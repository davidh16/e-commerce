package server

import (
	"e-commerce/config"
	"e-commerce/services"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"net/http"
)

func returnResponse(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
}

type Server struct {
	config  config.Config
	service services.Service
	db      *gorm.DB
	router  *mux.Router
}

func NewServer(s services.Service, cfg config.Config, postgres *gorm.DB, r *mux.Router) Server {

	server := Server{
		config:  cfg,
		service: s,
		db:      postgres,
		router:  r,
	}

	server.InitRoutes()

	return server
}

func (s *Server) InitRoutes() {
	// user routes
	s.router.HandleFunc("/register", s.Register).Methods("POST")
	s.router.HandleFunc("/login", s.Login).Methods("POST")
}
