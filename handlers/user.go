package handlers

import (
	"e-commerce/models"
	"encoding/json"
	"net/http"
)

//func RegisterUser(w http.ResponseWriter, req *http.Request) {
//	var user models.User
//	err := json.NewDecoder(req.Body).Decode(&user)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusBadRequest)
//		return
//	}
//
//}
//
//func LoginUser(w http.ResponseWriter, req *http.Request) {
//
//}
//
//func LogoutUser(w http.ResponseWriter, req *http.Request) {
//
//}
//
//func ResetUserPassword(w http.ResponseWriter, req *http.Request) {
//
//}

func (h *UserHandler) Register(w http.ResponseWriter, req *http.Request) {
	var user models.User
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	_, err = h.s.Create(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (h UserHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
}
