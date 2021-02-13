package golangbasics

import (
	"fmt"
	"time"
)

func Concurrency() {
	go say("There. ")
	say("Hi! ")
}

func say(s string) {
	i := 0
	for i < 3 {
		fmt.Println(s)
		time.Sleep(time.Millisecond * 100)
		i = i + 1
	}
}
