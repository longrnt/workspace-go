package main

import "fmt"

func main() {
	m := make([]string,1,25)
	fmt.Println(m)//[]
	changeMe(m)
	fmt.Println(m)//Todd

	k := 5
	fmt.Println(k)
	changePrimitive(k)
	fmt.Println(k)

}

func changeMe(z []string) {
	z[0] = "Todd"
	fmt.Println(z)//Todd
}

func changePrimitive(t int) {
	t = 10
	fmt.Println(t)
}