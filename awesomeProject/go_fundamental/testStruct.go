package main

import (
	"fmt"
	"time"
)

type Session struct {
	playerId  string
	beehive   string
	timestamp time.Time
}

func main() {
	session := Session{}
	if (Session{}) == session {
		fmt.Println("zero")
	}
	session1 := Session{playerId: "123"}

	fmt.Println(session1.playerId)
	fmt.Println(session.playerId)
}
