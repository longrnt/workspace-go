package main

import "fmt"

func main() {

	var n, _ = fmt.Println("Hello world version 2", "created by Long", 123)

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i, " ")
		}
	}

	fmt.Println(n)
	//fmt.Println(e)

}
