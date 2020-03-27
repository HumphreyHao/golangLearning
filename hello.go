package main

import "fmt"

const englishHelloPrefix = "Hello, "

// Hello ...
// Solutation : here you tell us what Salutation is
// Printer : what is this?
// Greet : describe what this function does
// CreateMessage : describe what this function does
func Hello(name string) string {
	if name == "" {
		name = "world"
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("happy"))
}
