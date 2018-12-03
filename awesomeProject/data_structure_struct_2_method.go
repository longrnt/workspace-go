package main

import "fmt"

type Person1 struct {
	firstName string
	lastName string
	age int
}

type DoubleZero1 struct {
	Person1
	first string
	licenseToKill bool
}

func (p Person1) fullName() {
	fmt.Println(p.firstName + " " + p.lastName)
}

func main() {
	person1:= DoubleZero1{
		Person1: Person1{
			firstName:"Long",
			lastName:"Tran",
			age:30,
		},
		first:"double zero seven",
		licenseToKill:true,
	}

	fmt.Println(person1)
	person1.fullName()
}
