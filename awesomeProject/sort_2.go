package main

import (
	"fmt"
	"sort"
)

/*
Use https://golang.org/pkg/sort/ to sort the following:

(1)
type people []string
studyGroup := people{"Zeno", "John", "Al", "Jenny"}

(2)
s := []string{"Zeno", "John", "Al", "Jenny"}

(3)
n := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}

Also sort the above in reverse order
 */


func main() {
	//(1)
	//type people []string
	//studyGroup := people{"Zeno", "John", "Al", "Jenny"}
	//fmt.Println(studyGroup)

	//sort.Strings(sort.Reverse(sort.StringSlice{studyGroup}))
	//fmt.Println(studyGroup)

	//(2)
	s := []string{"Zeno", "John", "Al", "Jenny"}
	fmt.Println(s)

	s = sort.StringSlice(s)
	//sort.Sort(sort.Reverse(sort.StringSlice(s)))

	i := sort.Reverse(sort.StringSlice(s))
	fmt.Println(i)

	sort.Sort(i)
	fmt.Println(i)

	//(3)
	n := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}
	fmt.Println(n)
	sort.Ints(n)
	fmt.Println(n)


}
