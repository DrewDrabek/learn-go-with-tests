package main

import "fmt"

const (
	spanish = "Spanish"
	french  = "French"

	englishHello = "Hello, "
	spanishHello = "Hola, "
	frenchHello  = "Bonjour, "
)

func greetingPrefix(language string) (prefix string) {
	switch language {
	case "Spanish":
		prefix = spanishHello
	case "French":
		prefix = frenchHello
	default:
		prefix = englishHello
	}
	return prefix
}

func Hello(user string, language string) string {

	if user == "" {
		user = "World"
	}

	return greetingPrefix(language) + user

}

func main() {
	fmt.Println(Hello("Drew", "Spanish"))
}
