// 日常用法;
package main

import "fmt"

func main() {
    book := make(map[string]string)

    book["price"] = "100"
    book["author"] = "wohu"
    book["language"] = "Chinese"

    for k, _ := range book { // 迭代 key, 不能保证每次迭代元素的顺序
        fmt.Println(k, "value is ", book[k])
    }

    for k, v := range book { // 同时迭代 key 和 value
        fmt.Println(k, v)
    }

    /*查看元素在集合中是否存在 */
    published, ok := book["published"] /*如果确定是真实的,则存在,否则不存在 */
    /*fmt.Println(ok) */
    if ok {
        fmt.Println("published is ", published)
    } else {
        fmt.Println("published not exist")
    }
}
