package main

import "fmt"

func main() {
	a:= 10
	var b *int = &a
	fmt.Println(a)
	fmt.Println(&a)


	fmt.Println(b)
	*b = 20
	fmt.Println(b)
	fmt.Println(*b)
	fmt.Println(a)

	i:=0
	for i<10 {
		fmt.Println(i)
		if i==5 {
			break
		}
		i++
	}
}
