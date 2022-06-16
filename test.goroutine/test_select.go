// 测试 time.after
// select case
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("开始时间：", time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println("time.Now()：", time.Now().Format("2006-01-02 15:04:05.999"))
	select {
	case <-time.After(time.Second * 2):
		fmt.Println("2秒后的时间：", time.Now().Format("2006-01-02 15:04:05"))
		fmt.Println("time.Now()：", time.Now())
	default:
		fmt.Println("default exit")
	}
}
