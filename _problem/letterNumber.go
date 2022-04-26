// 交替打印字母数字
// 12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718
// 要点: for *** select *** 结构  -- channel
package main

import (
    "fmt"
    "strings"
    "sync"
)

func main(){
    letter, number := make(chan bool), make(chan bool)
    wait := sync.WaitGroup{}

    go func() {
        i:=1
        for {
            select {
            case <-number:
                fmt.Print(i)
                i++
                fmt.Print(i)
                i++
                letter<-true
                break
            default:
                break
            }
        }

    }()

    wait.Add(1)
    go func(wait *sync.WaitGroup) {
        str := "ABCDEFGHIJKLMNOPQ"
        i := 0
        for {
            select {
            case <-letter:
                if i>= strings.Count(str, "")-2 {
                    wait.Done()
                    return
                }
                fmt.Print(str[i:i+1])
                i++
                fmt.Print(str[i:i+1])
                i++
                number <- true
                break
            default:
                break
            }
        }
    }(&wait)
    number<-true
    wait.Wait()
}
















