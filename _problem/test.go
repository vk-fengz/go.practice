// 手写数字字母交替打印
// 12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728
package main

import (
    "fmt"
    "strings"
    "sync"
)

func main() {
    letter, number := make(chan bool), make(chan bool)
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
        i := 0
        str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
        for i < strings.Count(str, "")-1 {
            <-letter
            fmt.Print(str[i : i+2])
            i += 2
            number <- true
        }
        wait.Done()
    }(&wait)

    number<-true
    wait.Wait()
}
