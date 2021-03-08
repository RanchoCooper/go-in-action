package main

import (
	"log"
	"os"

	_ "github.com/RanchoCooper/go-in-action/chapter02/matchers"
	"github.com/RanchoCooper/go-in-action/chapter02/search"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("president")
}
