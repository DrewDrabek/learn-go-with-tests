package main

import "fmt"

const englishHello = "Hello, "

func Hello(user string, language string) string {

	if user == "" {
		user = "World"
	}
	if language == "Spanish" {
		return "Hola, " + user
	}

	return englishHello + user

}

func main() {
	fmt.Println(Hello("Drew", "Spanish"))
}
