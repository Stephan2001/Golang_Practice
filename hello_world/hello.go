package main

import "fmt"

const englishprefix = "Hello, "

func main() {
	fmt.Println(Hello("world"))
}

func Hello(name string) string {
	if name == "" {
		name = "world"
	}
	return englishprefix + name
}
