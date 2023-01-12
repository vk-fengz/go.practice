package main

import (
	"sync"
)

var wg sync.WaitGroup

func main() {
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go work(&wg)
	}
	wg.Wait()
}

func work(wg *sync.WaitGroup) {
	var counter int
	for i := 0; i < 1e10; i++ {
		counter++
	}
	wg.Done()
}
