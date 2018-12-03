package main

import "fmt"

/*
Initialize a MAP using a composite literal where the key is a string and the value is an int;
print out the map;
range over the map printing out just the key;
range over the map printing out both the key and the value

 */
func main() {
	mapNumber := map[string]int {
		"One": 1,
		"Two": 2,
		"Three": 3,
	}

	fmt.Println(mapNumber)

	for k,_ := range mapNumber {
		fmt.Println(k)
	}

	for k,v := range mapNumber {
		fmt.Println(k,v)
	}
}
