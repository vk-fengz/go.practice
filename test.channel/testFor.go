// 非匿名函数
package main

import (
    "fmt"
    "time"
)

func getNumber(ch chan<- int) {
    // 从3循环到0
    for i := 3; i >= 0; i-- {
        // 发送3到0之间的数值
        ch <- i
        // 每次发送完时等待
        time.Sleep(time.Second)
    }
}
func main() {
    // 构建一个通道
    ch := make(chan int)
    // 开启一个并发匿名函数
    go getNumber(ch)

    // 遍历接收通道数据
    for data := range ch {
        // 打印通道数据
        fmt.Println(data)
        // 当遇到数据0时, 退出接收循环
        // 如果继续发送, 由于接收 goroutine 已经退出, 没有 goroutine 发送到通道, 因此运行时将会触发宕机报错;
        if data == 0 {
            break
        }
    }

}
