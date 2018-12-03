package main

import (
	"encoding/json"
	"fmt"
)

type Person3 struct {
	firstName string
	LastName string
	age int64
	notExported int64
}

type Person4 struct {
	First string
	Last string `json:"-"`
	Age int		`json:"tuoi"`
	notExported int
}

type Message struct {
	Name string
	Body string
	Time int64
}

func main() {
	p1:= Person3{"Long", "Tran", 30, 7}
	//p1:= Person3{firstName:"Long", lastName:"Tran", age:30, notExported:007}
	bs,_ := json.Marshal(p1)
	fmt.Println(bs)
	fmt.Printf("%T \n", bs)
	fmt.Println(string(bs))
	fmt.Println(p1)
	fmt.Println("-------------")
	//output,_ := json.Unmarshal(bs)
	//fmt.Println(output)

	m := Message{"Alice", "Hello", 1294706395881547000}
	b, _ := json.Marshal(m)
	fmt.Println(m)
	fmt.Println(b)
	fmt.Println(string(b))


	p2 := Person4{"Long", "Tran", 30, 007}
	b2,_ := json.Marshal(p2)
	fmt.Println(b2)
	fmt.Println(string(b2))
}
