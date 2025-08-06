package main

import "fmt"

const englishHello = "Hello, "

func Hello(user string) string {

	if user != "" {
		return englishHello + user + "!"
	} else {
		return "Hello, World"
	}
}

func main() {
	fmt.Println(Hello("Drew"))
}
