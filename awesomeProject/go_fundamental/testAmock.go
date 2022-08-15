package main

import "fmt"

type Performer interface {
	perform()
}

type A struct {
}

func (a *A) perform() {
	fmt.Println("real method")
}

type AMock struct {
}

func (a *AMock) perform() {
	fmt.Println("mocked method")
}

func caller(p Performer) {
	p.perform()
}

func main() {
	//test real method
	p := &A{}
	caller(p)

	//test mock method
	pMock := &AMock{}
	caller(pMock)

}
