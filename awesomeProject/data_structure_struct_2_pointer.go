package main

import "fmt"

type Person2 struct {
	firstName string
	lastName string
	age int
}

func main() {
	p1:= &Person2{"Long", "Tran", 30}
	fmt.Println(p1.firstName)
	fmt.Println(p1.lastName)
	fmt.Println(p1.age)

	fmt.Println(p1)
}
