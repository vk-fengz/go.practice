package main

import (
	"fmt"
	"sync/atomic"
)

var flag int64 = 0
var countt int64 = 0

func adda() {
	for {
		if atomic.CompareAndSwapInt64(&flag, 0, 1) {
			fmt.Println(flag)
			fmt.Println(countt)
			countt++
			atomic.StoreInt64(&flag, 0)

			return
		}
	}
}

func main() {
	go adda()
	go adda()
}
