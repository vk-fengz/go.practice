// 字符串 转换 rune类型
// unicode 字符的长度判断
package main

import "fmt"

func main() {
    var str string = "中国"
    rangeRune([]rune(str))
    rangeStr(str)
}

func rangeRune(arg []rune) {
    fmt.Println("rune type arg length is ", len(arg))
    for i := 0; i < len(arg); i++ {
        fmt.Printf("i is %d, value is %c\n", i, arg[i])
    }
}

func rangeStr(arg string) {
    fmt.Println("str type arg length is ", len(arg))
    for i := 0; i < len(arg); i++ {
        fmt.Printf("i is %d, value is %c\n", i, arg[i])
    }
}

