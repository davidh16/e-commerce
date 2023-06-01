package routes

import (
	"e-commerce/controller"
	"github.com/gorilla/mux"
)

func NewRouter(controller controller.Controller) *mux.Router {

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
