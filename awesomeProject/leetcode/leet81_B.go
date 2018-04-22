package main

import "fmt"

//Card Flipping Game
//https://leetcode.com/problems/card-flipping-game/description/
func main() {
	//fronts := []int{1,2,4,4,7}
	//backs := []int{1,3,4,1,3}
	fronts := []int{3,2}
	backs := []int{2,2}
	fmt.Println(flipgame(fronts, backs))
}

func flipgame(fronts []int, backs []int) int {
	var result = 2000000000

	var m = map[int]int{}
	var invalidMap = map[int]int{}

	for i:= 0; i < len(fronts); i++ {
		_, invalidFront := invalidMap[fronts[i]]
		_, invalidBack := invalidMap[backs[i]]
		if fronts[i] != backs[i] {
			if !invalidFront {
				m[fronts[i]] = fronts[i]
			}
			if !invalidBack {
				m[backs[i]] = backs[i]
			}
		} else {
			delete(m, fronts[i])
			invalidMap[fronts[i]] = fronts[i]
		}
	}
	for key, _ := range m {
		if key < result {
			result = key
		}
	}
	if result == 2000000000 {
		result = 0
	}
	return result
}