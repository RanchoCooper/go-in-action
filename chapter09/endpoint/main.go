package main

import (
	"github.com/RanchoCooper/go-in-action/chapter09/endpoint/handlers"
	"log"
	"net/http"
)

const PORT = ":4040"

func main() {
	handlers.Routes()

	log.Println("started listening on: ", PORT)
	_ = http.ListenAndServe(PORT, nil)
}
