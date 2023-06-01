package controller

import "net/http"

func (c Controller) HandleSomething() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		return
	}
}
