package users

import (
	"github.com/biFebriansyah/goback/src/middleware"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func New(rt *mux.Router, db *gorm.DB) {
	route := rt.PathPrefix("/users").Subrouter()

	repo := NewRepo(db)
	svc := NewService(repo)
	ctrl := NewCtrl(svc)

	route.HandleFunc("", middleware.Do(ctrl.GetByUsername, middleware.CheckAuth)).Methods("GET")
	route.HandleFunc("/all", middleware.Do(ctrl.GetAll, middleware.CheckAuth)).Methods("GET")
	route.HandleFunc("", ctrl.AddData).Methods("POST")
}
