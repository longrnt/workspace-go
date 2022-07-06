package main

import "fmt"

//https://leetcode.com/problems/climbing-stairs/
func main() {
	fmt.Println(climbStairs(4))
}

func climbStairs(n int) int {
	a := 1
	b := 1
	c := 0
	if n == 1 {
		return 1
	} else {
		for i := 2; i <= n; i++ {
			c = a + b
			a = b
			b = c
		}
	}
	return b
}
