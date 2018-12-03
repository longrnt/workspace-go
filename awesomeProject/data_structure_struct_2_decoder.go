package main

import (
	"strings"
	"encoding/json"
	"fmt"
)

type Person7 struct {
	First string
	Last string
	Age int
	notExported int
}

func main() {
	var p1 Person7
	rdr:= strings.NewReader(`{"First":"Long", "Last":"Tran", "Age":20}`)
	json.NewDecoder(rdr).Decode(&p1)

	fmt.Println(p1)
	fmt.Println(p1.First)
	fmt.Println(p1.Last)
	fmt.Println(p1.Age)
}
