package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenschHelloPrefix = "lalala, "
const spanish = "Spanish"
const french = "French"

// Hello ...
// Solutation : here you tell us what Salutation is
// Printer : what is this?
// Greet : describe what this function does
// CreateMessage : describe what this function does
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenschHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("happy", ""))
}
