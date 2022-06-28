package auth

import (
	"github.com/biFebriansyah/goback/src/modules/v1/users"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func New(rt *mux.Router, db *gorm.DB) {
	route := rt.PathPrefix("/auth").Subrouter()

	repo := users.NewRepo(db)
	svc := NewService(repo)
	ctrl := NewCtrk(svc)

	route.HandleFunc("", ctrl.Sigin).Methods("POST")
}
