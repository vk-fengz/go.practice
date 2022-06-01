// https://dev.to/mcaci/how-to-use-the-context-done-method-in-go-22me
// context.WithCancel(...) sends a value to the channel ctx.Done() that will end the for loop and exit.

package main

import (
	"context"
	"log"
	"time"
)

const interval = 500

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(5 * interval * time.Millisecond)
		cancel()
	}()
	f(ctx)
}

func f(ctx context.Context) {
	ticker := time.NewTicker(interval * time.Millisecond)
	for {
		select {
		case <-ticker.C:
			doSomething()
		case <-ctx.Done():
			return
		}
	}
}

func doSomething() { log.Println("tick") }
