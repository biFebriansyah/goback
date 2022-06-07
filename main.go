package main

import (
	"log"
	"net/http"

	"github.com/biFebriansyah/goback/src/routers"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	mainRoute, err := routers.New()
	if err != nil {
		log.Fatal(err.Error())
	}

	if err := http.ListenAndServe(":8080", mainRoute); err != nil {
		log.Fatal("aplikasi gagal dijalankan")
	}

}
