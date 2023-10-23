package server

import (
	"e-commerce/config"
	"e-commerce/services"
	"encoding/json"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"net/http"
)

func returnResponse(w http.ResponseWriter, status int, err error, custom any) {
	w.WriteHeader(status)
	if err != nil {
		w.Write([]byte(err.Error()))
	}

	if custom != nil {
		err = json.NewEncoder(w).Encode(custom)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
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

func (s *Server) InitRoutes() {
	// user routes
	s.router.HandleFunc("/register", s.Register).Methods("POST")
	s.router.HandleFunc("/login", s.Login).Methods("POST")

	// category routes
	s.router.HandleFunc("/category", s.CreateCategory).Methods("POST")
	s.router.HandleFunc("/category/update", s.UpdateCategory).Methods("POST")
	s.router.HandleFunc("/category/delete", s.DeleteCategory).Methods("POST")
	s.router.HandleFunc("/category", s.GetCategory).Methods("GET")
	s.router.HandleFunc("/category/list", s.ListCategories).Methods("GET")

	// subcategory routes
	s.router.HandleFunc("/subcategory", s.CreateSubcategory).Methods("POST")
	s.router.HandleFunc("/subcategory/update", s.UpdateSubcategory).Methods("POST")
	s.router.HandleFunc("/subcategory/delete", s.DeleteSubcategory).Methods("POST")
	s.router.HandleFunc("/subcategory", s.GetSubcategory).Methods("GET")
	s.router.HandleFunc("/subcategory/list", s.ListSubategories).Methods("GET")

	// media routes
	s.router.HandleFunc("/media/upload", s.CreateAndUploadMedia).Methods("POST")
	s.router.HandleFunc("/media/download", s.DownloadMedia).Methods("POST")

	// product routes
	s.router.HandleFunc("/product", s.CreateProduct).Methods("POST")
	s.router.HandleFunc("/product/update", s.UpdateProduct).Methods("POST")
	s.router.HandleFunc("/product/delete", s.DeleteProduct).Methods("POST")
	s.router.HandleFunc("/product", s.GetProduct).Methods("GET")
	s.router.HandleFunc("/product/list", s.ListProducts).Methods("GET")

	// address routes
	s.router.HandleFunc("/address", s.SaveAddress).Methods("POST")
	s.router.HandleFunc("/address", s.GetAddress).Methods("Get")
	s.router.HandleFunc("/address/list", s.ListUsersAddresses).Methods("GET")
}
