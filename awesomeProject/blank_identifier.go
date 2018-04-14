package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func main() {
	fmt.Print()
	res, error  := http.Get("http://youtube1.com/")
	if res != nil {
		page, _ := ioutil.ReadAll(res.Body)
		res.Body.Close()
		fmt.Printf("%s", page)
	} else {
		fmt.Println(error)
	}

}
