package main

import "fmt"

func half(n int) (float64, bool) {
	return float64(n)/2, n%2 == 0
}

func main() {
	n := 11
	m, h := half(n)
	fmt.Println(m, h)
}
