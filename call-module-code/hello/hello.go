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

	message, err := greetings.Hello("")

	if err != nil {
		// Will exit
		log.Fatal(err)
	}

	fmt.Println(message)
}
