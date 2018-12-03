package main

import "fmt"

type person struct {
	firstName string
	lastName string
	age int
}

func main() {
	long:= person{"long", "tran", 30}
	huy:= person{"huy", "do", 32}
	fmt.Println(long)
	fmt.Println(huy)
}
