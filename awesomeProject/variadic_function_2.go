package main

import "fmt"

func main() {
	numbers:= []int{1,2,3,4,5}
	print(numbers...)
}

func print(numbers ...int) {
	fmt.Println(numbers)
}