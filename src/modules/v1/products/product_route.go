package products

import (
	"github.com/biFebriansyah/goback/src/middleware"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func New(rt *mux.Router, db *gorm.DB) {
	route := rt.PathPrefix("/product").Subrouter()

	repo := NewRepo(db)
	ctrl := NewCtrl(repo)

	route.HandleFunc("", ctrl.GetAll).Methods("GET")
	route.HandleFunc("", middleware.Do(ctrl.AddData, middleware.CheckAuth)).Methods("POST")
}
