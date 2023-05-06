package handlers

import (
	"e-commerce/models"
	"encoding/json"
	"net/http"
)

func (h *UserHandler) Register(w http.ResponseWriter, req *http.Request) {
	var user models.User
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.s.Create(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

// this function redirects request to certain handler method, it is necessary to extend it if new handler method is introduced
func (h UserHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	switch path {
	case "/register":
		if req.Method == http.MethodPost {
			h.Register(w, req)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	default:
		http.Error(w, "Not found", http.StatusNotFound)
	}
}
