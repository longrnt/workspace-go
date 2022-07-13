package main

import "fmt"

type person1 struct {
	first string
	last  string
}

func (p person1) speak() {
	fmt.Println("I'm a person", p.first, p.last)
}

type secretAgent struct {
	person1
	ltk bool
}

func (sa secretAgent) speak() {
	fmt.Println("I'm a secret agent", sa.first, sa.last)
}

type human interface {
	speak()
}

func bar(h human) {

	switch h.(type) {
	case person1:
		fmt.Println("I was passed by a person1", h.(person1).first, h.(person1).last)
	case secretAgent:
		fmt.Println("I was passed by a secretAgent", h.(secretAgent).first, h.(secretAgent).last)
	}
	//fmt.Println("I was passed into bar", h)
}

func main() {
	sa1 := secretAgent{
		person1: person1{
			"James",
			"Bond",
		},
		ltk: true,
	}
	sa2 := secretAgent{
		person1: person1{
			"Miss",
			"Moneypenny",
		},
		ltk: true,
	}

	p1 := person1{
		"John",
		"Doe",
	}

	sa1.speak()
	sa2.speak()

	bar(p1)
	bar(sa1)
	bar(sa2)

}
