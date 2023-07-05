package routes

import (
	db2 "e-commerce/db"
	"e-commerce/models"
	"e-commerce/server"
	"encoding/json"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"net/http"
)

func NewRouter(controller server.Controller) *mux.Router {

	// setup of all routes that are available

	// example : [go setupNewModelRoutes(r, &corespondingHandler)]

	r := mux.NewRouter()
	go setupProductRoutes(r, controller)
	go setupCategoryRoutes(r, controller)
	go setupCustomerRoutes(r, controller)
	go setupOrderRoutes(r, controller)
	go setupUserRoutes(r, controller)
	return r
}

type MyRouter struct {
	R  *mux.Router
	db *gorm.DB
}

func NewMyRouter(db *gorm.DB, router *mux.Router) MyRouter {
	return MyRouter{
		db: db,
		R:  router,
	}
}
func (m MyRouter) Init() {
	m.R.HandleFunc("/xxx", m.RegisterTest)
	m.R.HandleFunc("/direct", m.DirectTest)
}

func (m MyRouter) Db() *gorm.DB {
	return m.db
}

func (m MyRouter) RegisterTest(w http.ResponseWriter, req *http.Request) {

	var user models.User

	// decoding json message to user model
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		return
	}

	x := m.db.Create(&user)

	if x.Error != nil {
		return
	}

	//TODO implement sending verification email

	return
}

func (m MyRouter) DirectTest(w http.ResponseWriter, req *http.Request) {
	var user models.User

	db := db2.ConnectToDb()
	// decoding json message to user model
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		return
	}

	x := db.Create(&user)

	if x.Error != nil {
		return
	}

	//TODO implement sending verification email

	return
}
