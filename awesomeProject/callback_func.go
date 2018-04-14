package main

import "fmt"

func calculate(numbers []int, callback func(int)) {
	for _, v := range numbers {
		callback(v)
	}
}

func main() {
	numbers := []int{1,2,3,4,5}
	calculate(numbers, func(v int) {
		fmt.Println(v)
	})
}
