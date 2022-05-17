package main

import (
	"fmt"
	"reflect"
)

type Dark interface {
	Run(string) string
	Speak(string) string
}
type bird struct {
	Name string
}

func (b *bird) Run(msg string) string {
	return msg
}
func (b *bird) Speak(string) string {
	return "asd"
}

var (
	// 判断结构体bird是否实现了Dark接口
	_ Dark = (*bird)(nil) // 把nil转化成*bird类型 赋值后即丢弃
)

func main() {
	var bird Dark = (*bird)(nil)

	fmt.Println(reflect.TypeOf(bird))
	fmt.Println("bird: ", bird.Run("ok"))
	fmt.Println("bird: ", bird.Speak("ok"))
}

// go 中 var _ Interface = &TestStruct{} 判断 结构体是否实现 所有方法
// 此方法用于 判断 结构体是否实现了 接口的 所有方法
// 此方法 在编译的 时候 就会检测 防止 出现调用出错
// var bird Dark = (*bird)(nil)
// bird 表示 结构体
// Dard 表示 接口
