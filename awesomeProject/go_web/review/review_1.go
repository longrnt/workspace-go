package main

import "fmt"

type hotdog int

type person struct {
	fname string
	lname string
}

type secretAgent struct {
	person
	licenseToKill bool
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func (p person) speak() {
	fmt.Println(p.fname, p.lname, `says, "Good morning, James."`)
}

func (sa secretAgent) speak() {
	fmt.Println(sa.fname, sa.lname, `says, "Shaken, not stirred."`)
}

func main() {
	xi := []int{2,1,4,5,6,8}
	fmt.Println(xi)

	m := map[string]int {
		"Todd": 45,
		"Job": 42,
	}

	fmt.Println(m)

	p1 := person{
		"Miss",
		"Moneypenny",
	}

	fmt.Println(p1.fname)
	p1.speak()

	pa1 := secretAgent{person{"Todd", "McClod"}, true}

	fmt.Println(pa1)

	pa1.speak()
	pa1.person.speak()

	saySomething(p1)
	saySomething(pa1)
}
