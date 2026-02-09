package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	// Disable time, source file and line number
	log.SetFlags(0)

	names := []string{"Gladys", "Samantha", "Darrin"}

	messages, err := greetings.Hellos(names)

	if err != nil {
		// Will exit
		log.Fatal(err)
	}

	fmt.Println(messages)
}
