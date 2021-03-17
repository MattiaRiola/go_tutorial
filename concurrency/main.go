package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//	sequentialStuff()
	// basicConcurrency()
	// syncConcurrency()
	channelConcurrency()
}

func advancedProcess(item string, out chan string) {
	// in this way i'm sure that I close the channel at the end of the process
	defer close(out)
	for i := 1; i <= 5; i++ {
		time.Sleep(time.Second)
		out <- item
	}
}
func advancedConsume(input chan string) {

	for msg := range input {
		//it will read from input channel untill it is closed
		fmt.Println(msg)
	}
	// for i := 1; i <= 5; i++ {
	// 	time.Sleep(time.Second)
	// 	fmt.Println("ready to receive block ", i, time.Now())
	// 	msg := <-input
	// 	fmt.Println("Block received")
	// 	fmt.Println(msg)
	// }
}

func process(item string) {
	for i := 1; i <= 5; i++ {
		time.Sleep(time.Second / 2)
		fmt.Println("Processed", i, item)
	}
}

func channelConcurrency() {
	out1 := make(chan string)
	go advancedProcess("order", out1)
	go advancedConsume(out1)
	fmt.Println("Press enter to stop")
	fmt.Scanln()
}

func syncConcurrency() {
	var wg sync.WaitGroup
	fmt.Println("--- Waiting group example ---")
	wg.Add(2)
	go func() {
		process("order")
		wg.Done()
	}()
	go func() {
		process("refund")
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("I've finished")
}

func basicConcurrency() {
	fmt.Println("I'm using go routine in order to handle concurrency")
	go process("order")
	go process("refund")
	fmt.Scanln()

}

func sequentialStuff() {
	fmt.Println("in this way i can only order because I'm using a single process")
	process("order")
	process("refund")
}
