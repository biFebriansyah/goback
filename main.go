package main

import (
	"log"
	"os"

	"github.com/biFebriansyah/goback/src/configs/command"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	if err := command.Run(os.Args[1:]); err != nil {
		log.Fatal(err)
	}
}
