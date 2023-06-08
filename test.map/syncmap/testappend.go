package main

import (
	"fmt"
	"sync"
	"time"
)

var sm sync.Map
var wg sync.WaitGroup

func main() {
	sl := []int{}
	// sl = append(sl, 0)
	sm.Store("aa", sl)

	// 读
	v, ok := sm.Load("aa")
	fmt.Println(len(v.([]int)))
	fmt.Printf("Load:\n v= %v\n ok=%v\n", v, ok)

	wg.Add(2)

	// func1
	go func(*sync.Map) {
		defer wg.Done()
		for i := 1; i <= 5000; i++ {
			if v, ok := sm.Load("aa"); ok {
				v = append(v.([]int), i)
				sm.Store("aa", v)
			}

		}
	}(&sm)

	// func2
	go func(*sync.Map) {
		defer wg.Done()
		time.Sleep(10 * time.Microsecond)
		for i := 10000; i <= 15000; i++ {
			if v, ok := sm.Load("aa"); ok {
				v = append(v.([]int), i)
				sm.Store("aa", v)
			}

		}
	}(&sm)

	// ======================================= 打印输出
	wg.Wait()
	v, ok = sm.Load("aa")
	fmt.Println("\n=====>  结果输出")
	fmt.Println(len(v.([]int)))
	// fmt.Printf("Load:\n slice= %v\n ok=%v\n", v, ok)

}

// func getmap(sm sync.Map) []int {

// 	v, ok := sm.Load("aa")
// 	if ok {
// 		return v
// 	}
// 	return nil
// }
