package main

import (
	"os"
	"runtime/trace"
)

func main() {
	trace.Start(os.Stderr)
	defer trace.Stop()

	ch := make(chan string)
	go func() {
		ch <- "Go 语言编程之旅"
	}()

	<-ch
}

// 参考: > https://golang2.eddycjy.com/posts/ch6/03-trace/
