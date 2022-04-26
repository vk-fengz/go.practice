package main

import "fmt"

func main() {
    a := 10
    defer func(i int) {
        fmt.Println("defer func i is ", i) // defer func i is  10
    }(a)

    a += 10
    fmt.Println("after defer a is ", a) // after defer a is  20

    // func1
    ret := func1()
    fmt.Println("func1 return value is  ", ret)
}

func func1() (i int) {
    defer func() {
        i++
        fmt.Println("defer2:", i)
    }()      // 声明匿名函数并直接执行
    return 0 // 结果为 1
}
