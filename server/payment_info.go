package server

import (
	"e-commerce/models"
	"encoding/json"
	"net/http"
)

func (s *Server) SavePaymentInfo(w http.ResponseWriter, req *http.Request) {
	var paymentInfo *models.PaymentInfo

	err := json.NewDecoder(req.Body).Decode(&paymentInfo)
	if err != nil {
		returnResponse(w, http.StatusBadRequest, err, nil)
		return
	}

	_, err = s.service.CreatePaymentInfo(paymentInfo)
	if err != nil {
		returnResponse(w, http.StatusInternalServerError, err, nil)
		return
	}

	returnResponse(w, http.StatusOK, nil, nil)
	return
}

func (s *Server) GetPaymentInfo(w http.ResponseWriter, req *http.Request) {
	var request struct {
		PaymentInfoUuid string `json:"paymentInfo_uuid"`
	}

	err := json.NewDecoder(req.Body).Decode(&request)
	if err != nil {
		returnResponse(w, http.StatusBadRequest, err, nil)
		return
	}

	paymentInfo, err := s.service.FindPaymentInfoByUuid(request.PaymentInfoUuid)
	if err != nil {
		returnResponse(w, http.StatusInternalServerError, err, nil)
		return
	}

	returnResponse(w, http.StatusOK, err, paymentInfo)
	return
}

func (s *Server) ListUsersPaymentInfo(w http.ResponseWriter, req *http.Request) {
	var request struct {
		UserUuid string `json:"user_uuid"`
	}

	err := json.NewDecoder(req.Body).Decode(&request)
	if err != nil {
		returnResponse(w, http.StatusBadRequest, err, nil)
		return
	}

	paymentInfo, err := s.service.FindPaymentInfoByUserUuid(request.UserUuid)
	if err != nil {
		returnResponse(w, http.StatusInternalServerError, err, nil)
		return
	}

	returnResponse(w, http.StatusOK, err, paymentInfo)
	return
}
