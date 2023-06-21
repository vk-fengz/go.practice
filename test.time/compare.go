package main

import (
	"fmt"
	"time"
)

func main() {
	// 用 p 简写
	p := fmt.Println

	// 当前时间
	now := time.Now()
	p("当前时间：", now) // 当前时间： 2022-04-12 16:58:13.1294251 +0800 CST m=+0.005672001

	weekday := time.Now().Weekday()
	p("当前是星期几:", weekday)
	// print next Week
	i := time.Now().Weekday()
	for j := 0; j < 7; j++ {
		if i == 7 {
			i = i - 7
		}
		p("range weekday:", int(i), i)
		i++
	}

}
