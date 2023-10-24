package main

import (
	"fmt"
	"time"
)

func handler(a, b int) string {
	t0 := time.Now()
	defer func() {
		fmt.Println("1. user time:", time.Since(t0).Milliseconds())
	}()
	if a > b {
		time.Sleep(100 * time.Millisecond)
		return "ok"
	} else {
		time.Sleep(200 * time.Millisecond)
		return "ok"
	}
}
func main() {

	handler(3, 4)
}
