package main

import "fmt"

func main() {
	x := []int{1, 4, 3, 5, 2}
	fmt.Println(x)

	x = append(x, 12)
	fmt.Println(x)

	fmt.Println(x[1:3])

	y := []int{222, 33, 444, 555, 333}

	x = append(x, y...)
	fmt.Println(x)
}
