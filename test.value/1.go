// 值传递;
// 指针同样是值传递, 拷贝指针的指向;
package main

import "fmt"

func main() {
    s := "脑子进煎鱼了"
    fmt.Println()
    fmt.Printf("main 内存地址：%p\n", &s)
    hello(&s)
    fmt.Printf("main 内存地址：%p\n", &s)
    fmt.Println(s)
}

func hello(s *string) {
    fmt.Printf("hello 内存地址：%p\n", &s)
    fmt.Println(*s)
    *s = "煎鱼进脑子了"

}
