package main

import "fmt"

func main() {
	x := sum(1, 2, 3, 4, 5)
	fmt.Println(x)
}

func sum(x ...int) int {
	var s int
	for _, v := range x {
		s += v
	}
	return s
}
