package main

import "fmt"

func main() {
	var mymap = map[int]string{
		1: "one",
		2: "two",
		3: "three",
	}
	println(mymap)
	fmt.Println(mymap)

	delete(mymap, 2)
	fmt.Println(mymap)
	mymap[5] = "five"
	fmt.Println(mymap)
	mymap[15] = "fifteen"
	fmt.Println(mymap)
}
