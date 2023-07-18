package main

import (
	"fmt"
	"time"
)

//类型定义
type NewInt int

//类型别名
type MyInt = int

func main() {
	var now = time.Now()

	fmt.Println(now)

	var a NewInt
	var b MyInt

	fmt.Printf("type of a:%T\n", a) //type of a:main.NewInt
	fmt.Printf("type of b:%T\n", b) //type of b:int

}
