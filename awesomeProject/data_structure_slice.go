package main

import "fmt"

func main() {
	//make(T, length, capacity)
	//sli := make([]int, 10, 50)
	mysclice := []string{"a", "b", "c", "g", "m", "z"}
	fmt.Println(mysclice)
	fmt.Println(mysclice[2:4]) 	// slicing a slice
	fmt.Println(mysclice[2])	// index access; accessing by index
	fmt.Println("myString"[2])	// index access; accessing by index

}
