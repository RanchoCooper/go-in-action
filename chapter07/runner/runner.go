// package runner manages the running and lifetime of a process
package runner

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

// Runner runs a set of tasks within a given timeout
// and can be shout down if an operating system interrupt occur
type Runner struct {
	interrupt chan os.Signal

	complete chan error

	timeout <- chan time.Time

	tasks []func(int)
}

var TimeoutErr = errors.New("received times")
var InterruptErr = errors.New("received interrupt")

// New returns a new ready-to-use Runner
func New(d time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		timeout:   time.After(d),
	}
}

func (r *Runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

// Start runs all tasks and monitors channel events.
func (r *Runner) Start() error {
	signal.Notify(r.interrupt, os.Interrupt)

	go func() {
		r.complete <- r.run()
	}()

	select {
	case err := <- r.complete:
		return err
	case <-r.timeout:
		return TimeoutErr
	}
}

// run executes each registered task.
func (r *Runner) run() error {
	for id, task := range r.tasks {
		if r.gotInterrupt() {
			return InterruptErr
		}
		task(id)
	}
	return nil
}

// gotInterrupt verifies if the interrupt signal has been issued
func (r *Runner) gotInterrupt() bool {
	select {
	case <-r.interrupt:
		signal.Stop(r.interrupt)
		return true
	default:
		return false
	}
}