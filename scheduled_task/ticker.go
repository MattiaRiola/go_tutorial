package main

import (
	"fmt"
	"os"
	"time"
)

func Ticker_solution(ok chan bool) {
	Task(time.Now())
	duration := time.Duration(10) * time.Second
	tk := time.NewTicker(duration)
	done := make(chan bool)

	go interrupt_ticker(done)
	for {
		select {
		case <-done:
			ok <- true
			return
		case t := <-tk.C:
			Task(t)
		}
	}

}

func interrupt_ticker(done chan bool) {

	keepasking := true
	for keepasking {
		fmt.Println("enter 'stop' to stop the ticker")
		var input string
		_, err := fmt.Scanln(&input)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		if input == "stop" {
			done <- true
			fmt.Println("ticker stopped")
			keepasking = false
		}
	}
}
