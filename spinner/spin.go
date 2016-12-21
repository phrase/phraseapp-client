package spinner

import (
	"fmt"
	"time"
)

func Spin(msg string, task func(finished chan<- struct{})) {
	// start task
	finished := make(chan struct{})
	go task(finished)

	fmt.Print(msg)

	// start spinning animation
	stop := make(chan struct{})
	go spin(stop)

	// wait for task to finish, then tell spinner to stop
	stop <- <-finished
	// wait for spinner to stop
	<-stop

	fmt.Println()
	return
}

// spin animates a spinner until it receives something on the stop channel. It then clears the spinning character and closes the stop channel, signaling that it's done.
func spin(stop chan struct{}) {
	chars := []string{`-`, `\`, `|`, `/`}
	fmt.Print(" ")
	i := 0
	for {
		fmt.Print("\b")
		fmt.Print(chars[i])
		select {
		case <-stop:
			fmt.Print("\b ")
			close(stop)
			return
		case <-time.After(100 * time.Millisecond):
		}

		i = (i + 1) % len(chars)
	}
}
