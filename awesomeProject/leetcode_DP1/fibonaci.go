package main

import "fmt"

//https://leetcode.com/problems/fibonacci-number/
func main() {
	fmt.Println(fib(0))
}

func fib(n int) int {
	a := 1
	b := 1
	c := 0
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		for i := 2; i < n; i++ {
			c = a + b
			a = b
			b = c
		}
	}
	return b
}
