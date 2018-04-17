package main

import "fmt"

func main() {
	var myslice = []int{}
	fmt.Println(myslice)
	fmt.Println(myslice == nil)

	myslice = append(myslice, 0)
	myslice = append(myslice, 1)
	myslice = append(myslice, 2)
	myslice = append(myslice, 3)
	fmt.Println(myslice)

	myslice = append(myslice[2:])
	fmt.Println(myslice)
}
