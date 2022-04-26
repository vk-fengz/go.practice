// 反射判断 interface 是否是nil值;
// 指针的判断;
package main

import (
    "fmt"
    "reflect"
)

func main() {
    var data *byte
    var in interface{}

    in = data
    fmt.Println(IsNil(in))
}

func IsNil(i interface{}) bool {
    vi := reflect.ValueOf(i)
    if vi.Kind() == reflect.Ptr {
        fmt.Println(vi)
        return vi.IsNil()
    }
    return false
}
