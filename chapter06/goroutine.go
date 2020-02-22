package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	// allocate 1 logical processor for the scheduler to use
	runtime.GOMAXPROCS(1)

	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("start goroutines")

	go func() {
		defer wg.Done()

		for count:= 0; count < 3; count++ {
			for char := 'a'; char < 'a' + 26; char++ {
				fmt.Printf("%c", char)
			}
		}
	}()

	go func() {
		defer wg.Done()

		for count:= 0; count < 3; count++ {
			for char := 'A'; char < 'A' + 26; char++ {
				fmt.Printf("%c", char)
			}
		}
	}()

	fmt.Println("waiting to finish")
	wg.Wait()
	fmt.Println("exit...")
}
