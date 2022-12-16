package greetings

import (
	"errors"
	"fmt"
)

// Toto returns a greeting for the named person
func Toto(name string) (string, error) {

	// If no name, return err message
	if name == "" {
		return "", errors.New("Empty name")
	}

	message := fmt.Sprintf("Hi, %v, welcome !", name)
	return message, nil
}
