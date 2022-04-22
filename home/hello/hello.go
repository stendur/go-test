package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("Steve")
	fmt.Println(message)
}
