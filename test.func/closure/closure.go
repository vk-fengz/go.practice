package main

import (
	"fmt"
)

// 累加器生成函数，这个函数输出一个初始值，
// 调用时返回一个为初始值创建的闭包函数

func Accumulate(value int) func() int {
	// 返回一个闭包,每次返回会创建一个新的函数实例
	return func() int {
		// 累加,对引用的 Accumulate 参数变量进行累加，注意 value 不是上一行匿名函数定义的，但是被这个匿名函数引用，所以形成闭包。
		value++
		// 返回一个累加值
		return value
	}
}

func main() {
	// 创建一个累加器，初始值为 1，返回的 accumulator 是类型为 func()int 的函数变量。
	accumulator := Accumulate(1)
	// 累加1并打印
	fmt.Println(accumulator())
	fmt.Println(accumulator())
	// 打印累加器的函数地址
	fmt.Printf("%p\n", &accumulator)

	// 创建一个累加器, 初始值为10
	accumulator2 := Accumulate(10)
	// 累加1并打印
	fmt.Println(accumulator2())
	// 打印累加器的函数地址
	fmt.Printf("%p\n", &accumulator2)
}
