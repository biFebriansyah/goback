package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func subRouter() {
	mainRoute := mux.NewRouter()

	mainRoute.HandleFunc("/", sampleHandler).Methods("GET")

	// localhost:8080/users {get, post, put}
	users := mainRoute.PathPrefix("/users").Subrouter()
	users.HandleFunc("/", usersHandler)

	// localhost:8080/product {get, post, put}
	product := mainRoute.PathPrefix("/product").Subrouter()
	product.HandleFunc("/", prodHandler)

	if err := http.ListenAndServe(":8080", mainRoute); err != nil {
		log.Fatal("aplikasi gagal dijalankan")
	}
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello from users"))
}

func prodHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello from prod"))
}
