// http://c.biancheng.net/view/4307.html

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan struct{}, 2)
	var l sync.Mutex
	go func() {
		l.Lock()
		defer l.Unlock()
		fmt.Println("goroutine1: 我会锁定大概 2s")
		time.Sleep(time.Second * 2)
		fmt.Println("goroutine1: 我解锁了，你们去抢吧")
		ch <- struct{}{}
	}()
	go func() {
		fmt.Println("goroutine2: 等待解锁")
		l.Lock()
		defer l.Unlock()
		fmt.Println("goroutine2: 欧耶，我也解锁了")
		ch <- struct{}{}
	}()
	// 等待 goroutine 执行结束
	for i := 0; i < 2; i++ {
		<-ch
	}
}

// goroutine2: 等待解锁
// goroutine2: 欧耶，我也解锁了
// goroutine1: 我会锁定大概 2s
// goroutine1: 我解锁了，你们去抢吧

// 需要注意的是一个互斥锁只能同时被一个 goroutine 锁定，其它 goroutine 将阻塞直到互斥锁被解锁（重新争抢对互斥锁的锁定）
