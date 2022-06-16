// 测试退出,
// 函数调用的影响
// rbmerger tolocal处理其他的   go post

package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	handle(ctx, "handle")
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("main: %d\n", i)
	}

	cancel()

	fmt.Println("===> exit main()")
}

func handle(ctx context.Context, str string) {

	go say(ctx, "goroutine: hello world")

	for i := 0; i < 5; i++ {
		// time.Sleep(10 * time.Millisecond)
		fmt.Printf("handle(): %d\n", i)
	}
	fmt.Println("===> exit handle()")
}

func say(ctx context.Context, s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(300 * time.Millisecond)
		fmt.Println(s, fmt.Sprintf(": %d", i))
	}
}
