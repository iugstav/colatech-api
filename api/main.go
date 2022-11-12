package main

import (
	"log"

	"github.com/iugstav/colatech-api/cmd/server"
)

func main() {
	err := server.InitServer()
	if err != nil {
		log.Fatal(err)
	}
}
