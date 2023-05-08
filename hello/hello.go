package main

import (
	"fmt"
	"log"

	greet "example.com/greetings" // alias
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greet.Hello("Luo")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
