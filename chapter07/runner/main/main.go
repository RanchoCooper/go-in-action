package main

import (
	"log"
	"os"
	"time"

	"github.com/RanchoCooper/go-in-action/chapter07/runner"
)

const timeout = 3 * time.Second

func main() {
	log.Println("Starting work.")

	// Create a new timer value for this run.
	r := runner.New(timeout)

	// Add the tasks to be run.
	tasks := []func(int) {
		createTask(),createTask(),createTask(),createTask(),
	}
	r.Add(tasks...)

	// Run the tasks and handle the result.
	if err := r.Start(); err != nil {
		switch err {
		case runner.TimeoutErr:
			log.Println("Terminating due to timeout.")
			os.Exit(1)
		case runner.InterruptErr:
			log.Println("Terminating due to interrupt.")
			os.Exit(2)
		}
	}

	log.Println("Process ended.")
}

// createTask returns an example task that sleeps for the specified
// number of seconds based on the id.
func createTask() func(int) {
	return func(id int) {
		log.Printf("Processor - Task #%d.", id)
		time.Sleep(time.Duration(id) * time.Second)
	}
}
