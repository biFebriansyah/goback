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

	route.HandleFunc("", middleware.Adapt(ctrl.GetUser, middleware.AuthWithRole("users")).ServeHTTP).Methods("GET")
	route.HandleFunc("/all", middleware.Adapt(ctrl.GetAll, middleware.AuthWithRole("admin")).ServeHTTP).Methods("GET")
	route.HandleFunc("", ctrl.AddData).Methods("POST")
}
