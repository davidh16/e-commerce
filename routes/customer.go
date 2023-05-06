package routes

import (
	"github.com/gorilla/mux"
	"sync"
)

func setupCustomerRoutes(r *mux.Router, wg *sync.WaitGroup) {
	defer wg.Done()
}
