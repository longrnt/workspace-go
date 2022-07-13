package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := person{
		first: "Long",
		last:  "Tran",
		age:   34,
	}

	fmt.Println(p1)

	p2 := struct {
		make, model string
		year        int
	}{
		make:  "Toyota",
		model: "Camry",
		year:  2011,
	}

	fmt.Println(p2)
}