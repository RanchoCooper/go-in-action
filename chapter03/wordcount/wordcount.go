package main

import (
	"fmt"
	"github.com/RanchoCooper/go-in-action/chapter03/words"
	"io/ioutil"
	"os"
)

// main is the entry point for the application
func main() {
	fmt.Println(os.Args[0])
	filename := os.Args[1]

	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("error when open file:", err)
		return
	}

	text := string(contents)
	count := words.CountWords(text)
	fmt.Printf("There are %d words in your text. \n", count)
}
