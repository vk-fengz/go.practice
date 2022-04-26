// for + go func 输出顺序

package main

import (
    "fmt"
    "runtime"
    "sync"
)

func main() {
    runtime.GOMAXPROCS(1)
    wg := sync.WaitGroup{}

    wg.Add(10)

    for i := 0; i < 5; i++ {
        go func(i int) {
            fmt.Println("i: ", i)
            wg.Done()
        }(i)
    }


    for j := 0; j < 5; j++ {
        go func(j int) {
            fmt.Println("j : ", j)
            wg.Done()
        }(j)
    }

    wg.Wait()
}

// 结果输出:
// j :  4
// i:  0
// i:  1
// i:  2
// i:  3
// i:  4
// j :  0
// j :  1
// j :  2
// j :  3
