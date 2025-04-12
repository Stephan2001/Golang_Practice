package main

import "fmt"

const englishprefix = "Hello, "
const spanishprefix = "Hola, "

func main() {
	fmt.Println(Hello("world", englishprefix))
}

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}
	if language == "spanish" {
		return spanishprefix + name
	}
	return englishprefix + name
}
