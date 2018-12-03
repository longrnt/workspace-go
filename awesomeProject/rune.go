package main

import "fmt"

func main() {
	//fmt.Println([]byte("hello"))
	for i := 50; i <= 105; i++ {
		fmt.Println(i, " - ", string(i), " - ", []byte(string(i)))
	}

	fmt.Println( ('i'))
}
