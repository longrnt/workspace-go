package main

import "fmt"
//Shortest Distance to a Character
//https://leetcode.com/contest/weekly-contest-81/problems/shortest-distance-to-a-character/

func main() {
	//var S = "loveleetcode"
	//var S = "eloveleetcode"
	var S = "aaba"

	var C byte = 'b'
	var result = shortestToChar(S, C)
	fmt.Println(result)
}

func shortestToChar(S string, C byte) []int {
	result:= make([]int, len(S))
	for i:= 0; i < len(S); i++ {
		result[i] = 2000000000
	}

	i:= 0
	for i < len(S) {
		for i < len(S){
			if S[i] == C {
				break
			}
			i++
		}
		if i == len(S) {
			break
		}
		var count = 0
		for j:= i; j >= 0; j-- {
			if result[j] > count {
				result[j] = count
			} else {
				break
			}
			count++
		}
		count = 1
		for j:= i+1; j < len(S); j++ {
			if result[j] > count {
				result[j] = count
			} else if S[j] == C {
				break
			}
			count++
		}
		i++
	}
	return result
}