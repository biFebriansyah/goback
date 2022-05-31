package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type foods struct {
	Name  string
	Price int
}

func handlerReqs() {

	mainRoute := mux.NewRouter()

	mainRoute.HandleFunc("/", sampleHandler).Methods("GET")
	mainRoute.HandleFunc("/users/{id}", paramsHandler).Methods("GET")
	mainRoute.HandleFunc("/product", queryHandler).Methods("GET")
	mainRoute.HandleFunc("/food", bodyHandler).Methods("POST")

	if err := http.ListenAndServe(":8080", mainRoute); err != nil {
		log.Fatal("aplikasi gagal dijalankan")
	}
}

func sampleHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello worlds"))
}

func paramsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, "Params : %v", vars["id"])
}

func queryHandler(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	fmt.Fprintf(w, "Params : %v", vars["name"][0])
}

func bodyHandler(w http.ResponseWriter, r *http.Request) {
	var food foods

	json.NewDecoder(r.Body).Decode(&food)

	fmt.Println(food.Name)

	// fmt.Fprintf(w, "Params : %v", vars["name"][0])
}
