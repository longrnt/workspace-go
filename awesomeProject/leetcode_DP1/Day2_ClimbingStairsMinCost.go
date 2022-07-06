package main

import (
	"fmt"
)

//https://leetcode.com/problems/min-cost-climbing-stairs/
//Say f[i] is the final cost to climb to the top from step i. Then f[i] = cost[i] + min(f[i+1], f[i+2]).

func main() {
	cost := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	fmt.Println(minCostClimbingStairs(cost))
}

func minCostClimbingStairs(cost []int) int {
	//var f []int
	cost = append(cost, 0)

	for i := len(cost) - 3; i >= 0; i-- {
		cost[i] += min(cost[i+1], cost[i+2])
	}

	return min(cost[0], cost[1])
}

func min(a int, b int) int {
	if a <= b {
		return a
	}
	return b
}
