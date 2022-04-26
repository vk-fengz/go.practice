// 空结构体访问方法;
package main

import "fmt"

func main() {
    var m *man
    fmt.Println(m.GetName())
}

type man struct {
}

func (m *man) GetName() string {
    return "asong"
}

// 运行结果
// asong
