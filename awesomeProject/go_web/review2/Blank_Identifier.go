package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func main() {
	fmt.Println("Hello world")
	res,error := http.Get("http://youtube1.com")
	if res != nil {
		page, _ := ioutil.ReadAll(res.Body)
		res.Body.Close()
		fmt.Println("%s", page)
	} else {
		fmt.Println(error)
	}
}
