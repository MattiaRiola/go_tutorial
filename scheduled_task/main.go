package main

import (
	"fmt"
	"log"
	"time"
)

func Task(t time.Time) {
	fmt.Println("[main Task] time: ", time.Now().UTC(), "\t param time: ", t)
}

var DEBUG = true

func main() {
	fmt.Println("Hello world!")
	var mode string
	if DEBUG {
		mode = "cron"
	} else {
		fmt.Println("Enter the mode of the scheduler: [interval/cron]")
		_, err := fmt.Scanln(&mode)
		if err != nil {
			log.Fatal("error reading stdin", err)
		}
	}

	switch mode {
	case "interval":
		ok1 := make(chan bool)
		go Ticker_solution(ok1)
		if <-ok1 {
			fmt.Println("All ok")
		} else {
			fmt.Println("something not ok")
		}
	case "cron":
		Cron_solution()
	}

	fmt.Println("byebye!")
}
