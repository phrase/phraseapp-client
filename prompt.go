package main

import (
	"bufio"
	"fmt"
	"os"
)

var stdin = bufio.NewReader(os.Stdin)

// prompt prints msg, then reads a line of user input. The input line is then scanned into the args using fmt.Sscan().
//
// This doesn't use fmt.Scanln() because prompt() is often called in a loop (running until user input is valid)
// and Scanln returns two seperate errors for example when scanning into one integer and "a\n" is read from stdin,
// resulting in the prompt message being printed twice.
func prompt(msg string, args ...interface{}) error {
	fmt.Print(msg)

	line, err := stdin.ReadString('\n')
	if err != nil {
		return err
	}

	_, err = fmt.Sscan(line, args...)
	return err
}

func promptWithDefault(prompt string, arg *string, defaultValue string) error {
	fmt.Print(prompt)
	fmt.Printf("[default %v] ", defaultValue)

	n, err := fmt.Scanln(arg)
	if n != 1 {
		*arg = defaultValue
		return nil
	}

	return err
}
