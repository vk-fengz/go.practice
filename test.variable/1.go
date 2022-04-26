package main

import "fmt"

// 测试初始化;
// const 声明多个变量;
func main() {
	const (
		a, b = "golang", 100
		d, e
		f bool = true
		g		// 多变量声明
	)
	fmt.Println(d, e, g)
}
