package main

import (
	_ "github.com/RanchoCooper/go-in-action/chapter02/matchers"
	"github.com/RanchoCooper/go-in-action/chapter02/search"
	"log"
	"os"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("president")
}
