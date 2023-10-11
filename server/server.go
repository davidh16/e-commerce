package server

import (
	"context"
	"e-commerce/config"
	"e-commerce/services"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"net/http"
)

func returnResponse(ctx context.Context, w http.ResponseWriter, status int, err error) {
	tx := ctx.Value("tx")

	if status != 200 {
		tx.(*gorm.DB).Rollback()
	} else {
		tx.(*gorm.DB).Commit()
	}

	w.WriteHeader(status)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
}

type Server struct {
	config  config.Config
	service services.Service
	router  *mux.Router
}

func NewServer(s services.Service, cfg config.Config, postgres *gorm.DB, r *mux.Router) Server {

	server := Server{
		config:  cfg,
		service: s,
		router:  r,
	}

	server.InitRoutes()

	return server
}

type Response struct {
	Status  int
	Message string
	Error   string
}

func (s *Server) InitRoutes() {
	// user routes
	s.router.HandleFunc("/register", s.Register).Methods("POST")
	s.router.HandleFunc("/login", s.Login).Methods("POST")
}
