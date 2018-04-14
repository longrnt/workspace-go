package main

import "fmt"

func makeGreetFunc() func() string {
	return func() string {
		return "Hello"
	}
}

func main() {
	greetFunc := makeGreetFunc()
	fmt.Println(greetFunc())
	fmt.Printf("%T", greetFunc)
}
