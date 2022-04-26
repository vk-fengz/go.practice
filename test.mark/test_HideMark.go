// 展示无法从另一个包中访问未公开的标志符
package main

import (
    "code.practise/test.mark/counters"
    "fmt"
)

func main() {
    // 创建一个未公开类型的变量
    // 并初始化为10
    // counter := counters.alterCounter(10)
    counter := counters.New(10)

    fmt.Printf("Counter: %d", counter)

}
