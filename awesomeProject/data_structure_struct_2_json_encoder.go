package main

import (
	"encoding/json"
	"os"
)

type Person6 struct {
	First string
	Last string
	Age int
	notExported int
}
func main() {
	p1:= Person6{"Long", "Tran", 20, 007}
	json.NewEncoder(os.Stdout).Encode(p1)
}
