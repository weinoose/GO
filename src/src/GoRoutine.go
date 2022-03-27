package src

import (
	"fmt"
	"time"
)

func Concurrency() {
	go say("There.")
	say("Hi!")
}

func say(s string) {
	i := 0
	for i < 3 {
		fmt.Println(s)
		time.Sleep(time.Millisecond * 1000)
		i = i + 1
	}
}

func pinger(c chan string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}

func printer(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func Channel() {
	var c chan string = make(chan string)

	go pinger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)
}
