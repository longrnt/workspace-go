package main

import "fmt"

type Person struct {
	firstName string
	lastName string
	age int
}

type DoubleZero struct {
	Person
	first string
	licenseToKill bool
}


func main() {
	person1:= DoubleZero{
		Person: Person{
			firstName:"Long",
			lastName:"Tran",
			age:30,
		},
		first:"double zero seven",
		licenseToKill:true,
	}

	fmt.Println(person1)
}
