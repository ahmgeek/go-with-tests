package main

import "fmt"

var greetings = map[string]string {
	"": "Hello, ",
	"Spanish": "Hola, ",
	"Frensh": "Bonjour, ",
}

const helloPrefix = "Hello, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	
	return greetings[language] + name
}

func main() {
	fmt.Printf(Hello("World", ""))
}
