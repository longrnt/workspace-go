package main

import "fmt"

const (
	byte = iota
	//KB = 1 << (iota*10)
	KB = iota
	//MB = 1 << (iota*10)
	MB = iota
	//GB = 1 << (iota*10)
	GB = iota
)

func main() {
	fmt.Println("binary\t\tdecimal")
	//fmt.Printf("%b\t", KB)
	//fmt.Printf("%d\t", KB)
	//fmt.Printf("%b\t", MB)
	//fmt.Printf("%d\t", MB)
	//fmt.Printf("%b\t", GB)
	//fmt.Printf("%d\t", GB)


	fmt.Println(byte)
	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(GB)

}
