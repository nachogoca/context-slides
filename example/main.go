package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	//background()
	withCancel()
	//withTimeOut()
}

func background() {
	ctx := context.Background()

	sleepAndTalk(ctx, 3*time.Second, "Hello background!")
}

func withCancel() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	go func() {
		s := bufio.NewScanner(os.Stdin)
		s.Scan()
		cancel()
	}()

	sleepAndTalk(ctx, 3*time.Second, "Hello withCancel!")
}

func withTimeOut() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	sleepAndTalk(ctx, 5*time.Second, "Hello withTimeOut!")
}

func sleepAndTalk(ctx context.Context, d time.Duration, s string) {
	select {
	case <-time.After(d):
		fmt.Println(s)
	case <-ctx.Done():
		log.Println(ctx.Err())
	}
}
