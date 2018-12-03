package main

import "fmt"

type animal struct {
	sound string
}

type dog struct {
	animal
	friendly bool
}

type cat struct {
	animal
	annoying bool
}

func specs(a interface{}) {
	fmt.Println(a)
}

func main() {
	fido := dog{animal{"woof"}, true}
	fifi := cat{animal{"meow"},true}
	specs(fido)
	specs(fifi)

	y := make([]int, 0, 10)
	y = append(y, 888)
	for i, v := range y {
		fmt.Println(i, " - ", v)
	}
}
