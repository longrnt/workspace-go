package main

import "fmt"

func main() {
	fmt.Println("This is the most awesome program in the world")

	foo()

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Print(i)
		}
	}

}

func foo() {
	fmt.Println("This is foo")
}
