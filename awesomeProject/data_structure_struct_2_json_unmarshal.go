package main

import (
	"encoding/json"
	"fmt"
)

type Person5 struct {
	First string
	Last string `json:"-"`
	Age int		`json:"tuoi"`
	notExported int
}

func main() {
	var p1 Person5
	bs:= []byte(`{"First":"Long","Last":"Tran","tuoi":30}`)
	json.Unmarshal(bs, &p1)

	fmt.Println(p1)

}
