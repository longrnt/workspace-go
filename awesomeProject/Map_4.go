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

	val, exist := mymap[1]
	fmt.Println(val, exist)

	val, exist = mymap[5]
	fmt.Println(val, exist)
	fmt.Println(val == "")

	if v,e:= mymap[3]; e {
		fmt.Println(v)
		fmt.Println(e)
	} else {
		fmt.Println("value not exist")
	}
}
