// 结构体和指针调用
package main

import "fmt"

type Vertex struct {
	Name1 string
	Name2 string
}

func main() {
	v := Vertex{"脑子进了", "煎鱼"}
	v.Name2 = "蒸鱼"
	fmt.Println(v.Name2)

	var vnil Vertex
	if vnil.Name1 == "" {
		fmt.Println("nil vnil Name1 access success")
	} else {
		fmt.Println("fail")
	}

}
