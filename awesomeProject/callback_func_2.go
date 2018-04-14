package main

import "fmt"

func visit(numbers []int, callback func(int) bool) []int {
	var result []int

	for _, v := range numbers {
		if callback(v) {
			result = append(result, v)
		}
	}
	return result
}

func main() {
	var xs = []int{1,2,3,4,5,6,7,8,9,10}
	result := visit(xs, func(v int) bool {
		return v > 5
	})
	fmt.Println(result)
}
