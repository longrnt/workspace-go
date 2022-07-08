package main

import "fmt"

func main() {
	x := []int{1, 4, 3, 5, 2}
	fmt.Println(x)

	x = append(x, 12)
	fmt.Println(x)
}
