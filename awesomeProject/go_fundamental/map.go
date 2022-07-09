package main

import "fmt"

func main() {
	m := map[string]int{
		"Long": 33,
		"Hue":  31,
	}

	fmt.Println(m)

	fmt.Println(m["Long"])

	if v, ok := m["Hue1"]; ok {
		fmt.Println("Key exist in map, value = ", v)
	} else {
		fmt.Println("Key doesn't exist")
	}

	m["Hang"] = 22

	for k, v := range m {
		fmt.Println(k, v)
	}
}
