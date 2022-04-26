// 字符串中最后子串的长度
// hello nowcoder
// 输出: 8
// https://www.nowcoder.com/practice/8c949ea5f36f422594b306a2300315da?tpId=37&&tqId=21224&rp=1&ru=/ta/huawei&qru=/ta/huawei/question-ranking
package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func LastStringLength(str *string) int {
    newStr := strings.Fields(*str) // 库函数
    return len(newStr[len(newStr)-1])
}

// 填 fmt.scanf 的坑  读取输入的字符串
func Scanf(input *string) {
    reader := bufio.NewReader(os.Stdin)
    data, _, _ := reader.ReadLine()
    *input = string(data)
}

func main() {
    var input string
    Scanf(&input)
    // fmt.Printf("%d", LastStringLength(&input))
    // 10年刷acm题的时候就是这种方法输出结果，10年了一点长进没有
    fmt.Printf("%d", LastStringLength1(&input))

}

// 手撕
func LastStringLength1(str *string) int {

    j := 0
    for i := len(*str) - 1; i >= 0; i-- {
        // if unicode.IsSpace(rune((*str)[i])) {
        //     return j
        // }
        if (*str)[i] == ' ' {
            return j
        }
        j++

    }
    return j
}
