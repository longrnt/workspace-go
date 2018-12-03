package main

import "fmt"

func main() {
	n:= 20
	m:= 20

	for i:= 0; i < n; i++ {
		for j:= 0; j < m; j++ {
			if i == 0 || i == n - 1 || j == 0 || j == m-1 || i == j || i == m-j-1 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}

		}
		fmt.Println()
	}
}
