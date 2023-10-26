package server

import (
	"e-commerce/models"
	"encoding/json"
	"errors"
	"github.com/samber/lo"
	"net/http"
)

func (s *Server) AddItemToTheCart(w http.ResponseWriter, req *http.Request) {
	var cartItem models.CartItem
	err := json.NewDecoder(req.Body).Decode(&cartItem)
	if err != nil {
		returnResponse(w, http.StatusBadRequest, err, nil)
		return
	}

	me, err := s.service.Me(req.Header.Get("Authorization"))
	if err != nil {
		returnResponse(w, http.StatusBadRequest, err, nil)
		return
	}

	cart, err := s.service.GetCart(me)
	if err != nil {
		returnResponse(w, http.StatusInternalServerError, err, nil)
		return
	}

	if cart == nil {
		cart.Uuid = "uuid.New().String()"
		cart.UserUuid = me
		cart.Items = append(cart.Items, cartItem)

		cart, err = s.service.AddItemToACart(cart)
		if err != nil {
			returnResponse(w, http.StatusInternalServerError, err, nil)
			return
		}

		returnResponse(w, http.StatusOK, nil, cart)
		return
	}

	updatedCart := lo.FilterMap(cart.Items, func(item models.CartItem, index int) (*models.Cart, bool) {
		if item.ProductUuid == cartItem.ProductUuid {
			cart.Items[index].Quantity += cartItem.Quantity
			return cart, true
		}
		return nil, false
	})[0]
	if updatedCart != nil {
		updatedCart, err = s.service.AddItemToACart(updatedCart)
		if err != nil {
			returnResponse(w, http.StatusInternalServerError, err, nil)
			return
		}
		returnResponse(w, http.StatusOK, nil, updatedCart)
		return
	}

	cart.Items = append(cart.Items, cartItem)
	cart, err = s.service.AddItemToACart(cart)
	if err != nil {
		returnResponse(w, http.StatusInternalServerError, err, nil)
		return
	}
	returnResponse(w, http.StatusOK, nil, cart)
	return
}

func (s *Server) RemoveItemFromTheCart(w http.ResponseWriter, req *http.Request) {
	var cartItem models.CartItem
	err := json.NewDecoder(req.Body).Decode(&cartItem)
	if err != nil {
		returnResponse(w, http.StatusBadRequest, err, nil)
		return
	}

	me, err := s.service.Me(req.Header.Get("Authorization"))
	if err != nil {
		returnResponse(w, http.StatusBadRequest, err, nil)
		return
	}

	cart, err := s.service.GetCart(me)
	if err != nil {
		returnResponse(w, http.StatusInternalServerError, err, nil)
		return
	}

	if cart == nil {
		returnResponse(w, http.StatusNotFound, errors.New("cart does not exist"), nil)
		return
	}

	//TODO ZAVRSITI
	updatedCart := lo.FilterMap(cart.Items, func(item models.CartItem, index int) (*models.Cart, bool) {
		if item.ProductUuid == cartItem.ProductUuid {
			if item.Quantity > 1 {
				cart.Items[index].Quantity -= 1
			} else {
				cart.Items = append(cart.Items[:index], cart.Items[index+1:]...)
			}
			return cart, true
		}
		return nil, false
	})[0]
	if updatedCart != nil {
		updatedCart, err = s.service.RemoveItemFromACart(updatedCart)
		if err != nil {
			returnResponse(w, http.StatusInternalServerError, err, nil)
			return
		}
		returnResponse(w, http.StatusOK, nil, updatedCart)
		return
	}

	cart.Items = append(cart.Items, cartItem)
	cart, err = s.service.RemoveItemFromACart(cart)
	if err != nil {
		returnResponse(w, http.StatusInternalServerError, err, nil)
		return
	}

	returnResponse(w, http.StatusOK, nil, cart)
	return
}

func (s *Server) GetCart(w http.ResponseWriter, req *http.Request) {
	cart, err := s.service.GetCart(req.FormValue("user_uuid"))
	if err != nil {
		returnResponse(w, http.StatusInternalServerError, err, nil)
		return
	}
	returnResponse(w, http.StatusOK, nil, cart)
	return
}
