package main

// should it be main or hello ?
// It should be main if we want to run this file directly
// It should be hello if we want to import this file as a package in another file

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("Greetings : ")
	names := []string{"Vighnesh", "Aryan", "Ashutosh"}

	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
}
