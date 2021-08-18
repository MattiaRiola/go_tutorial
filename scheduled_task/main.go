package main

import (
	"fmt"
	"time"
)

func Task(t time.Time) {
	fmt.Println("I'm running task, time: ", t)
}

func main() {
	fmt.Println("Hello world!")
	ok := make(chan bool)
	go Ticker_solution(ok)
	if <-ok {
		fmt.Println("All ok")
	} else {
		fmt.Println("something not ok")
	}
}
