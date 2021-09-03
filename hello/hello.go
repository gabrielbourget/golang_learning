package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("George")
	fmt.Println(message)
}
