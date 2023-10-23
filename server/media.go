package server

import (
	"context"
	"e-commerce/models"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

func (s *Server) CreateAndUploadMedia(w http.ResponseWriter, req *http.Request) {

	if req.FormValue("category_uuid") == "" || req.FormValue("subcategory_uuid") == "" || req.FormValue("product_uuid") == "" {
		returnResponse(w, http.StatusBadRequest, errors.New("missing data"), nil)
		return
	}

	path := strings.Join([]string{req.FormValue("category_uuid"), req.FormValue("subcategory_uuid"), req.FormValue("product_uuid")}, "/")

	file, header, err := req.FormFile("media")
	if err != nil {
		returnResponse(w, http.StatusBadRequest, err, nil)
		return
	}

	media := &models.Media{
		Size:     header.Size,
		Filename: header.Filename,
		Path:     path,
	}

	_, err = s.service.CreateAndUploadMedia(media, file)
	if err != nil {
		returnResponse(w, http.StatusBadRequest, err, nil)
	}

	returnResponse(w, http.StatusOK, nil, nil)
	return
}

func (s *Server) DownloadMedia(w http.ResponseWriter, req *http.Request) {
	var request struct {
		MediaUuid string `json:"media_uuid" validate:"required"`
	}

	err := json.NewDecoder(req.Body).Decode(&request)
	if err != nil {
		returnResponse(w, http.StatusBadRequest, err, nil)
		return
	}

	media, err := s.service.FindMediaByUuid(request.MediaUuid)
	if err != nil {
		returnResponse(w, http.StatusNotFound, nil, nil)
		return
	}

	mediaFile, err := s.service.DownloadMedia(context.Background(), media.Path)
	if err != nil {
		returnResponse(w, http.StatusInternalServerError, err, nil)
		return
	}
	if mediaFile == nil {
		returnResponse(w, http.StatusNotFound, nil, nil)
		return
	}

	response, err := json.Marshal(map[string][]byte{
		"content": mediaFile,
	})
	if err != nil {
		returnResponse(w, http.StatusInternalServerError, err, nil)
		return
	}

	returnResponse(w, http.StatusOK, err, response)
	return
}
