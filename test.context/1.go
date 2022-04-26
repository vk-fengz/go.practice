// context 控制 goroutine
package main

import (
    "context"
    "fmt"
    "sync"
    "time"
)

func main() {
    var wg sync.WaitGroup
    wg.Add(3)
    ctx, stop := context.WithCancel(context.Background())
    go func() {
        defer wg.Done()
        watchingDog(ctx, "watchdog_1")
    }()

    go func() {
        defer wg.Done()
        watchingDog(ctx, "watchdog_2")
    }()

    go func() {
        defer wg.Done()
        watchingDog(ctx, "watchdog_3")
    }()

    time.Sleep(5 * time.Second)
    stop() // 发停止指令
    wg.Wait()
}

func watchingDog(ctx context.Context, name string) {
    // 开启for select循环, 一直后台监控
    for {
        select {
        case <-ctx.Done():
            fmt.Println(name, "receive stop cmd, will stop")
            return
        default:
            fmt.Println(name, "is running ……")
        }
        time.Sleep(1 * time.Second)
    }
}
