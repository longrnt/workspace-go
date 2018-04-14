package main

import "fmt"

func average(sf ...float64) float64 {
	fmt.Println(sf)
	fmt.Printf("%T \n",sf)
	total:= 0.0
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}

func main() {
	n:= average('a',99)
	arr := []float64{1,3,4,5,6,7}
	m:= average(arr...)
	fmt.Println(n)
	fmt.Println(m)
}
