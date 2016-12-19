package main

import "fmt"

var errInsufficientInput = fmt.Errorf("not enough tokens to scan")

func prompt(prompt string, args ...interface{}) error {
	fmt.Print(prompt)
	n, err := fmt.Scanln(args...)
	if n < len(args) {
		return errInsufficientInput
	}
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
