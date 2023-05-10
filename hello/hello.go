package main

import (
	"fmt"
	"log"

	greet "example.com/greetings" // alias
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Luo", "Zhong", "Lin"}
	messages, err := greet.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
}
