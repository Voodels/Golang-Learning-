package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello returns a greeting message for the named person.

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("Empty Name")
	}
	format, err := randomFormat()
	if err != nil {
		return "", err
	}
	message := fmt.Sprintf(format, name)
	// nil - no error if there is no error nil will be empty if there is an error it will contain the error message
	return message, nil
}

func randomFormat() (string, error) {
	// return a random format from a set of formats
	formats := []string{
		"Hi %v, Welcome!",
		"Hello %v, How are you?",
		"Greetings %v!",
	}
	return formats[rand.Intn(len(formats))], nil
}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}
