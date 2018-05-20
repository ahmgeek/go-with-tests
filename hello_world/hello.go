package hello

import "fmt"

const spanish = "Spanish"
const french = "French"

const englishPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

var greetings = map[string]string {
	"": "Hello, ",
	"Spanish": "Hola, ",
	"Frensh": "Bonjour, ",
}

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	

	return greetingPrefix(language) + name
	
	// Using maps
	//return greetings[language] + name
}

func greetingPrefix(language string) (prefix string){
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishPrefix
	}

	return
}

func main() {
	fmt.Printf(Hello("World", ""))
}
