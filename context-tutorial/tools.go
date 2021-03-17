package main

import (
	"context"
	"fmt"
	"time"
)

func sleepAndTalk(ctx context.Context, d time.Duration, msg string) {
	time.Sleep(d)
	fmt.Println(msg)
}
