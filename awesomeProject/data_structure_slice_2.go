package main

import "fmt"

func main() {
	myslice := make([]int, 5, 10)
	myslice[0] = 0
	myslice[1] = 1
	myslice[2] = 2
	myslice[3] = 3
	myslice[4] = 4
	myslice = append(myslice, 5)
	myslice = append(myslice, 6)
	myslice = append(myslice, 7)
	myslice = append(myslice, 8)
	myslice = append(myslice, 9)
	fmt.Println(cap(myslice))
	fmt.Println(myslice)

	myslice = append(myslice, 10)
	fmt.Println(cap(myslice))
	fmt.Println(myslice)

	for i:=0; i<10; i++ {
		myslice = append(myslice, myslice[len(myslice)-1] +1)
	}

	fmt.Println(cap(myslice))
	fmt.Println(myslice)

}
