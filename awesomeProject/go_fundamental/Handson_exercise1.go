package main

import (
	"fmt"
	"net"
)

var x int = 42
var y string = `James Bond`
var z bool = true

func main() {
	fmt.Println(x, y, z)
	fmt.Printf("%v %v %v", x, y, z)

	_, err := net.Dial("tcp", "scanme.nmap.org:80")
	if err == nil {
		fmt.Println("Connection Successful")

	}
}
