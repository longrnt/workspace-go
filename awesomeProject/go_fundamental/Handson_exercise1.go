package main

import "fmt"

var x int = 42
var y string = `James Bond`
var z bool = true

func main() {
	fmt.Println(x, y, z)
	fmt.Printf("%v %v %v", x, y, z)
}
