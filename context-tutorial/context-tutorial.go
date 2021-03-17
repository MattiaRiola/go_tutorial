package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func timeoutSleepAndTalk(ctx context.Context, d time.Duration, msg string) {
	select { // with the select I can check the context with select
	case <-time.After(d): //when i'm able to read from that go inside this case
		fmt.Println(msg)
	case <-ctx.Done(): //if the context is canceled
		log.Print(ctx.Err())
	}
}
func main() {
	contextTimeout()

}

func contextTimeout() {

	ctx := context.Background()

	// I had a cancel option for that context
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	//With timeout I've to cancel when I've finished because the system allocate some resources for us
	defer cancel()
	timeoutSleepAndTalk(ctx, 9000*time.Second, "hello")
	fmt.Println("OSS: I've context deadline exceeded as error message because the timer expired")

	ctx1 := context.Background()
	ctx1, cancel1 := context.WithTimeout(ctx1, 5*time.Second)
	cancel1()
	timeoutSleepAndTalk(ctx1, 2*time.Second, "hello")

}

func contextCancel() {
	ctx := context.Background()
	// I had a cancel option for that context
	ctx, cancel := context.WithCancel(ctx)
	go func() {
		time.Sleep(1 * time.Second)
		cancel()
	}()
	sleepAndTalk(ctx, 3*time.Second, "hello")
}
