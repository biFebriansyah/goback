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

	route.HandleFunc("", middleware.Handle(ctrl.GetUser, middleware.AuthTest("users"))).Methods("GET")
	route.HandleFunc("/all", middleware.Handle(ctrl.GetAll, middleware.AuthTest("admin"))).Methods("GET")
	route.HandleFunc("", ctrl.AddData).Methods("POST")
}
