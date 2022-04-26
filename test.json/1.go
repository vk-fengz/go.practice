package main

import "fmt"
import "encoding/json"

type Stu struct {
    Name string `json:"stu_name"`
    ID   string `json:"stu_id"`
    Age  int    `json:"-"`      // - 的作用;
}

func main() {
    buf, _ := json.Marshal(Stu{"Tom", "t001", 18})
    fmt.Printf("%s\n", buf)
}