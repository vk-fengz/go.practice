// 交替打印字母 数字
// WaitGroup 实现两协程同步;
// 12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728
// 参考: https://zhuanlan.zhihu.com/p/374921491
package main

import (
    "fmt"
    "sync"
    "unicode/utf8"
)

func main() {
    number, letter := make(chan bool), make(chan bool)
    wait := sync.WaitGroup{}
    go func() {
        i := 1
        for {
            <-number
            fmt.Printf("%d%d", i, i+1)
            i += 2
            letter <- true
        }
    }()
    wait.Add(1)
    go func(wait *sync.WaitGroup) {
        str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
        i := 0
        for {
            <-letter
            if i >= utf8.RuneCountInString(str) {
                wait.Done()
                return
            }
            fmt.Print(str[i : i+2])
            i += 2
            number <- true
        }
    }(&wait)
    number <- true
    wait.Wait()
}


