// 测试 struct json的解码与编码
package main

import (
    "encoding/json"
    "fmt"
)

type People struct {
    Name string `json:"name"`
}

func main() {
    js := `{"name":"11"}`

    var p People
    err := json.Unmarshal([]byte(js), &p)
    if err != nil {
        fmt.Println("err: ", err)
        return
    }
    fmt.Println("People: ", p)
}
