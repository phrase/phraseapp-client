package spinner

import (
	"fmt"
	"time"
)

// While executes f, displays an animated spinner while f runs, and stops when f returns.
func While(f func()) {
	c := make(chan struct{})

	go func(c chan<- struct{}) {
		defer close(c)
		f()
	}(c)

	Until(c)
}

// Until displays an animated spinner until reading from c succeeds. It is recommended to simply close c to achieve this.
func Until(c <-chan struct{}) {
	// start spinning animation
	stop := make(chan struct{})
	go spin(stop)

	// wait for task to finish, then tell spinner to stop
	stop <- <-c
	// wait for spinner to stop
	<-stop
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
