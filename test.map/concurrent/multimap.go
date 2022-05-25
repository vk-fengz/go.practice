// struct 中加锁 解决map并发写入 fatal error: concurrent map writes
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type multimap struct {
	sync.RWMutex
	// sync.Mutex
	mapIndex map[int]string
	mapCont  map[string]string
}

// *sync.WaitGroup 需要传递， 否则all goroutines are asleep - deadlock!
func rwmap(wg *sync.WaitGroup, testmap *multimap, index int) {

	defer wg.Done()
	str := fmt.Sprintf("%d", index)

	testmap.Lock()
	// fmt.Println(index)
	// 两条语句可以 复现 fatal error: concurrent map writes
	testmap.mapIndex[index] = "aa"
	testmap.mapIndex[index+1] = "aa"
	testmap.mapCont[str] = str
	testmap.Unlock()
}

func main() {
	var wg sync.WaitGroup
	rand.Seed(time.Now().UnixNano())

	testmap := new(multimap)
	testmap.mapCont = make(map[string]string)
	testmap.mapIndex = make(map[int]string)

	// 计时
	startTime := time.Now()

	for i := 0; i < 100000; i++ {
		wg.Add(1)
		index := rand.Intn(100)

		go rwmap(&wg, testmap, index)

		// go func(int) {
		// 	defer wg.Done()
		// 	fmt.Println(index)
		// 	testmap.mapIndex[index] = "aa"
		// }(index)

	}

	wg.Wait()

	// 计时器
	elapsedTime := time.Since(startTime)
	fmt.Println("Segment finished in ", elapsedTime) // Segment finished in xxms

	fmt.Println(testmap)

}
