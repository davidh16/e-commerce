package routes

import (
	"github.com/gorilla/mux"
	"sync"
)

func NewRouter() *mux.Router {

	r := mux.NewRouter()

	g := sync.WaitGroup{}

	//for efficient code, it is necessary to define a number of go routines that are going to be performed
	g.Add(5)

	// setup of all routes that are available
	// if a new model along with its own routes is introduced, it is necessary to increment number of go routines and define setup function
	// example : [go setupNewModelRoutes(r, &g)]
	go setupProductRoutes(r, &g)
	go setupCategoryRoutes(r, &g)
	go setupCustomerRoutes(r, &g)
	go setupOrderRoutes(r, &g)
	go setupUserRoutes(r, &g)

	g.Wait()

	return r
}
