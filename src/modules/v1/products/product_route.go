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
	route.HandleFunc("/test", middleware.Handle(ctrl.TESTPOST, middleware.FileUpload)).Methods("POST")
	route.HandleFunc("/{id}", ctrl.GetById).Methods("GET")
	route.HandleFunc("", middleware.Handle(ctrl.AddData, middleware.AuthTest("users"), middleware.FileUpload)).Methods("POST")
}
