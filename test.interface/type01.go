// 类型断言等
package main

import (
    "fmt"
)

// var container = []string{"aaa", "bbb", "ccc"}

func main() {
    container := map[string]string{"a": "aaa", "b": "bbb", "c": "ccc"}

    // 方式1;类型断言
    _, ok1 := interface{}(container).([]string)
    _, ok2 := interface{}(container).(map[string]string)

    // %T 表示该值的 Go 类型
    if !(ok1 || ok2) {
        fmt.Printf("Error: unsupported container type: %T\n", container)
        return
    }
    fmt.Printf("The element is %#v , (container type: %T)\n", container["a"], container)

    // 方式2;
    elem, err := getElement(container)
    if err != nil {
        fmt.Printf("Error: %s\n", err)
        return
    }
    fmt.Printf("The element is %#v , (container type: %T)\n", elem, container)
}

//	空接口包含所有的类型, 输入的参数均会被转换为空接口
// 	函数入参已经声明 containerI 为 interface 类型, 所以不需要再次进行 interface{}(container) 转换
func getElement(containerI interface{}) (elem string, err error) {
    //	变量类型会被保存在t中
    // 	方式2;类型查询
    switch t := containerI.(type) {
    case []string:
        elem = t[1]
    case map[string]string:
        elem = t["a"]
    default:
        err = fmt.Errorf("unsupported container type: %T", containerI)
        return
    }
    return
}
