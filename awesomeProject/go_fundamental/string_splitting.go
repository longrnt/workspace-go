package main

import (
	"fmt"
	"strings"
)

func main() {
	//str := "?yo.streamId=NSBAPlusHD-1999951.dfw.1080.mobile.cps.hls.lp60.v410.mssn.42&yo.up="
	var str string
	arr := strings.Split(str, "=")
	arr1 := strings.Split(arr[1], "&")
	result := arr1[0]
	fmt.Println(str)
	fmt.Println(arr)
	fmt.Println(arr1)
	fmt.Println(result)

}
