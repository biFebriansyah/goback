package carts

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/biFebriansyah/goback/src/database/gorm/models"
	"github.com/biFebriansyah/goback/src/interfaces"
	"github.com/gorilla/mux"
)

type cart_ctrl struct {
	repo interfaces.CartService
}

func NewCtrl(rep interfaces.CartService) *cart_ctrl {
	return &cart_ctrl{rep}
}

func (rep *cart_ctrl) GetByUserId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	vars := mux.Vars(r)["id"]

	data, err := rep.repo.GetByUserId(vars)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	json.NewEncoder(w).Encode(data)
}

func (rep *cart_ctrl) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var data models.Cart

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}

	result, err := rep.repo.Save(&data)
	if err != nil {
		fmt.Fprint(w, err.Error())
	}

	result.Send(w)
}
