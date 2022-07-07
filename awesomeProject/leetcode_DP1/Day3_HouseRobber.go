package main

import "fmt"

func main() {
	nums := []int{2, 7, 9, 3, 1}
	fmt.Println(rob(nums))
}

func rob(nums []int) int {
	var rob1 int
	var rob2 int

	for i := 0; i < len(nums); i++ {
		temp := max(nums[i]+rob1, rob2)
		rob1 = rob2
		rob2 = temp
	}

	return rob2
}

func max(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}
