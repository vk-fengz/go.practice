// context.WithTimeout 时间控制 goroutine
package main

import (
    "fmt"
    "sync"
    "time"
    "context"
)

var (
    wg sync.WaitGroup
)

func startTask(ctx context.Context) error {
    defer wg.Done()

    for i := 0; i < 30; i++ {
        select {
        case <-time.After(1 * time.Second):
            fmt.Printf("in goroutine do task %v\n", i)

        // we received the signal of cancelation in this channel
        case <-ctx.Done():
            fmt.Printf("cancel goroutine task %v\n", i)
            return ctx.Err()
        }
    }
    return nil
}

func main() {
    timeoutCtx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
    defer cancel()  // control close

    fmt.Println("startTask")

    wg.Add(1)
    go startTask(timeoutCtx)
    wg.Wait()

    fmt.Println("endTask")
}
