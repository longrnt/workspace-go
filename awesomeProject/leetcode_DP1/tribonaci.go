package main

import "fmt"

//https://leetcode.com/problems/n-th-tribonacci-number/

func main() {
	fmt.Println(tribonacci(4))
}

func tribonacci(n int) int {
	a := 0
	b := 1
	c := 1
	d := 0
	if n == 0 {
		return 0
	} else if n == 1 || n == 2 {
		return 1
	} else {
		for i := 2; i < n; i++ {
			d = a + b + c
			a = b
			b = c
			c = d
		}
	}
	return d
}
