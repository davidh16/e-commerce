package server

import (
	"e-commerce/config"
	mw "e-commerce/middleware"
	"e-commerce/services"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	"log"
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

func NewServer(s services.Service, cfg config.Config, redis *redis.Client) Server {

	r := mux.NewRouter()

	middleware := mw.InitializeMiddleware(redis)

	server := Server{
		config:  cfg,
		service: s,
		router:  r,
	}

	server.InitRoutes(middleware)

	return server
}

func (s *Server) InitRoutes(mw *mw.Middleware) {
	// role routes
	s.router.HandleFunc("/role/create", mw.AdminAuthMiddleware(s.CreateRole)).Methods("POST")
	s.router.HandleFunc("/role/delete", mw.AdminAuthMiddleware(s.DeleteRole)).Methods("POST")
	s.router.HandleFunc("/roles/list", mw.AdminAuthMiddleware(s.ListRoles)).Methods("GET")

	// user routes
	s.router.HandleFunc("/register", s.Register).Methods("POST")
	s.router.HandleFunc("/login", s.Login).Methods("POST")
	s.router.HandleFunc("/logout", mw.AuthMiddleware(s.Logout)).Methods("POST")
	s.router.HandleFunc("/me", mw.AuthMiddleware(s.Me)).Methods("GET")
	s.router.HandleFunc("/verify", s.VerifyAccount).Methods("POST")

	// category routes
	s.router.HandleFunc("/category", mw.AdminAuthMiddleware(s.CreateCategory)).Methods("POST")
	s.router.HandleFunc("/category/update", mw.AdminAuthMiddleware(s.UpdateCategory)).Methods("POST")
	s.router.HandleFunc("/category/delete", mw.AdminAuthMiddleware(s.DeleteCategory)).Methods("POST")
	s.router.HandleFunc("/category", mw.AuthMiddleware(s.GetCategory)).Methods("GET")
	s.router.HandleFunc("/category/list", mw.AuthMiddleware(s.ListCategories)).Methods("GET")

	// subcategory routes
	s.router.HandleFunc("/subcategory", mw.AdminAuthMiddleware(s.CreateSubcategory)).Methods("POST")
	s.router.HandleFunc("/subcategory/update", mw.AdminAuthMiddleware(s.UpdateSubcategory)).Methods("POST")
	s.router.HandleFunc("/subcategory/delete", mw.AdminAuthMiddleware(s.DeleteSubcategory)).Methods("POST")
	s.router.HandleFunc("/subcategory", mw.AuthMiddleware(s.GetSubcategory)).Methods("GET")
	s.router.HandleFunc("/subcategory/list", mw.AuthMiddleware(s.ListSubategories)).Methods("GET")

	// media routes
	s.router.HandleFunc("/media/upload", mw.AdminAuthMiddleware(s.CreateAndUploadMedia)).Methods("POST")
	s.router.HandleFunc("/media/download", mw.AdminAuthMiddleware(s.DownloadMedia)).Methods("POST")

	// product routes
	s.router.HandleFunc("/product", mw.AdminAuthMiddleware(s.CreateProduct)).Methods("POST")
	s.router.HandleFunc("/product/update", mw.AdminAuthMiddleware(s.UpdateProduct)).Methods("POST")
	s.router.HandleFunc("/product/delete", mw.AdminAuthMiddleware(s.DeleteProduct)).Methods("POST")
	s.router.HandleFunc("/product", mw.AuthMiddleware(s.GetProduct)).Methods("GET")
	s.router.HandleFunc("/product/list", mw.AuthMiddleware(s.ListProducts)).Methods("GET")

	// address routes
	s.router.HandleFunc("/address", mw.AuthMiddleware(s.SaveAddress)).Methods("POST")
	s.router.HandleFunc("/address", mw.AuthMiddleware(s.GetAddress)).Methods("Get")
	s.router.HandleFunc("/address/list", mw.AuthMiddleware(s.ListUsersAddresses)).Methods("GET")

	// cart routes
	s.router.HandleFunc("/cart/add", mw.AuthMiddleware(s.AddItemToTheCart)).Methods("POST")
	s.router.HandleFunc("/cart", mw.AuthMiddleware(s.GetCart)).Methods("Get")
	s.router.HandleFunc("/cart/remove", mw.AuthMiddleware(s.RemoveItemFromTheCart)).Methods("POST")
}

func (s *Server) Serve() {
	err := http.ListenAndServe(s.config.Port, s.router)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("server running on port: ", s.config.Port)
}
