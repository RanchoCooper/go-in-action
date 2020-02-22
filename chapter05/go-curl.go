package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func init() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ./go-curl <url>")
		os.Exit(-1)
	}
}

func main() {
	res, err := http.Get(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	_, _ = io.Copy(os.Stdout, res.Body)
	if err := res.Body.Close(); err != nil {
		fmt.Println(err)
	}
	
}
