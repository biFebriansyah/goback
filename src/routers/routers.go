package routers

import (
	"net/http"

	"github.com/biFebriansyah/goback/src/configs/database"
	"github.com/biFebriansyah/goback/src/modules/v1/products"
	"github.com/gorilla/mux"
)

func New() (*mux.Router, error) {
	mainRoute := mux.NewRouter()

	db, err := database.New()
	if err != nil {
		return nil, err
	}

	mainRoute.HandleFunc("/", sampleHandler).Methods("GET")
	products.New(mainRoute, db)

	return mainRoute, nil
}

func sampleHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello worlds"))
}
