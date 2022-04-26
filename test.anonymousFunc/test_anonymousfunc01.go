// 匿名函数 与 flag

package main

import (
    "flag"
    "fmt"
)

// 定义命令行参数 skill, 从命令行输入 --skill 可以将 = 后的字符串传入 skillParam 指针变量;
var skillParam = flag.String("skill", "", "skill to perform")

func main() {
    // 解析命令行参数, 解析完成后, skillParam 指针变量将指向命令行传入的值;
    flag.Parse()
    // 定义一个从字符串映射到 func() 的 map, 然后填充这个 map
    var skill = map[string]func(){
        "fire": func() { // 初始化 map 的键值对, 值为匿名函数;
            fmt.Println("chicken fire")
        },
        "run": func() {
            fmt.Println("soldier run")
        },
        "fly": func() {
            fmt.Println("angel fly")
        },
    }
    /* skillParam 是一个 *string 类型的指针变量, 
       使用 *skillParam 获取到命令行传过来的值, 并在 map 中查找对应命令行参数指定的字符串的函数;
    */
    if f, ok := skill[*skillParam]; ok {
        f()
    } else {
        fmt.Println("skill not found")
    }

}
