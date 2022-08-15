package main

import "fmt"

func main() {

	x := foo1()
	fmt.Println(x)
	fmt.Println(x())

}

func foo1() func() int {
	return func() int {
		return 123
	}
}
