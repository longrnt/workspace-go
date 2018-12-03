package main

import "fmt"

/*
1. Initialize a SLICE of int using a composite literal;
print out the slice;
range over the slice printing out just the index;
range over the slice printing out both the index and the value
 */

func main() {
	arr := []int{1,3,4,6,7}
	fmt.Println(arr)

	for i,_ := range arr {
		fmt.Println(i)
	}
	for i,v := range arr {
		fmt.Println(i, v)
	}
}
