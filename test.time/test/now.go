package main

import (
	"fmt"
	"time"
)

func main() {
	// 用 p 简写
	p := fmt.Println

	// var 变量
	var tnow time.Time
	p("var变量时间:", tnow)
	p("var变量时间早于now:", tnow.Before(time.Now()))
	tnow = time.Now()
	p("var重新赋值now:", tnow)

	// 取地址
	aint := 10
	p("aint的地址:", &aint, aint)
	// 当前时间
	now := time.Now()
	p("当前时间：", now) // 当前时间： 2022-04-12 16:58:13.1294251 +0800 CST m=+0.005672001
	p("当前时间地址", &now)

	p(now.After(now))

	// time 运算
	sub := time.Now().Sub(tnow)
	p("时间差值:", sub)

}
