package main

import "fmt"

func main() {
	m := map[string][]int{
		"Long": {33, 44, 55},
		"Hue":  {31, 12},
		"Ben":  {11},
	}

	fmt.Println(m)

	fmt.Println(m["Long"])

	if v, ok := m["Hue"]; ok {
		fmt.Println("Key exist in map, value = ", v)

		v = append(v, 69)
		m["Hue"] = v

		fmt.Println("after adding _69_ = ", m["Hue"])
	} else {
		fmt.Println("Key doesn't exist")
	}

	if v, ok := m["Hue1"]; ok {
		fmt.Println("Key exist in map, value = ", v)
	} else {
		fmt.Println("Key doesn't exist")
	}

	//m = append(m, "")
	//m["Hang"] = {22, 44}
	//
	//for k, v := range m {
	//	fmt.Println(k, v)
	//}
}
