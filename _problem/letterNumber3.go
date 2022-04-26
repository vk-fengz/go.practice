// 交替打印字母数字: 使用单独channel 实现两协程同步;
// AB12CD34EF56GH78IJ910KL1112MN1314OP1516QR1718ST1920UV2122WX2324YZ2526
package main

import (
    "fmt"
    "runtime"
    "unicode/utf8"
)
func main() {
    runtime.GOMAXPROCS(runtime.NumCPU())
    number := make(chan bool)
    letter := make(chan bool)
    done := make(chan struct{})

    go func() {
        i := 1
        for {
            <-number
            fmt.Printf("%d%d", i, i+1)
            i += 2
            letter <- true
        }
    }()

    go func() {
        str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
        i := 0
        for {
            <-letter
            if i >= utf8.RuneCountInString(str) {
                break
            }
            fmt.Print(str[i : i+2])
            i += 2
            number <- true
        }
        done <- struct{}{}  // 空结构体 信号
    }()

    number <- true      // 控制先打印的一方;
    <-done
}

