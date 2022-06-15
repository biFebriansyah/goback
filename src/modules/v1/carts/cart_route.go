package carts

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func New(rt *mux.Router, db *gorm.DB) {
	route := rt.PathPrefix("/carts").Subrouter()

	repo := NewRepo(db)
	svc := NewService(repo)
	ctrl := NewCtrl(svc)

	route.HandleFunc("/{id}", ctrl.GetByUserId).Methods("GET")
	route.HandleFunc("/", ctrl.Create).Methods("POST")
}
