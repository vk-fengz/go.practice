// strings 常用操作
// 类型转换, 格式化等
package main

import (
    "fmt"
    "strconv"
    "strings"
)

func main() {
    value, err := strconv.ParseInt("33225", 10, 32)
    if err != nil {
        fmt.Println("error happens")
    }
    fmt.Println(value)

    flag, err := strconv.ParseBool(string('0'))
    if err != nil {
        fmt.Println("error happens")
    }
    fmt.Println(flag)

    fmt.Println(strconv.FormatBool(true))

    // 字符串运算
    str1 := "abcdefabcg"
    str2 :="abc"
    // cmp := strings.Compare(str1, str2)
    cmp := strings.Compare(str2, str1)
    fmt.Println("stringCompare:", cmp)

    var theInd = strings.Index(str1, "bc")
    repeat := strings.Repeat("abc", 10)

    fmt.Println(theInd)
    fmt.Println(repeat)
}
