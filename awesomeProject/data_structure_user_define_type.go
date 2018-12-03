package main

import "fmt"

type foo int

func main() {
	var myAge foo
	myAge = 30

	var yourAge int
	yourAge = 40

	//This shouldn't work
	//fmt.Println(myAge + yourAge)

	//This should work
	fmt.Println(int(myAge) + yourAge)

}
