package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	// 用 p 简写
	p := fmt.Println

	// 取地址
	aint := 10
	p("aint的地址:", &aint, aint)
	// 当前时间
	now := time.Now()
	p("当前时间：", now) // 当前时间： 2022-04-12 16:58:13.1294251 +0800 CST m=+0.005672001
	p("当前时间地址", &now)

	utcnow := time.Now().UTC()
	p("当前时间utc：", utcnow)

	// 获取年，其它如月，日等内容同理
	year := now.Year()
	p("当前年份：", year) // 前年份： 2022
	// 获取星期几
	weekday := now.Weekday()
	p("当前是周几：", weekday)          // 当前是周几： Tuesday
	p("当前是周几 pointer：", &weekday) // 当前是周几： Tuesday
	p("明天是周几：", now.Weekday()+1)  //

	// 构造时间，Date 的最后一个参数是时区
	myTime := time.Date(1996, time.April, 27, 0, 0, 0, 0, time.UTC)
	localTime := time.Date(1996, time.April, 27, 0, 0, 0, 0, time.Now().Location())
	p("我构造的时间：", myTime)
	p("我构造的时间：", localTime)

	// 时间比较，有 Before(), Equal(), After()
	// 这里展示 Before() 就都懂了
	// 首先构建两个时间
	my2000Year := time.Date(2000, time.April, 27, 0, 0, 0, 0, time.UTC)
	my3000Year := time.Date(3000, time.April, 27, 0, 0, 0, 0, time.UTC)

	is2000Before3000 := my2000Year.Before(my3000Year)
	p("2000 年在 3000 年之前吗：", is2000Before3000) // true

	// 时间计算，如加减，这里展示减法
	start := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2000, 1, 1, 12, 0, 0, 0, time.UTC)

	diff := end.Sub(start)
	p("end - start 的时间，结果是：", diff) // 当前时间减去 my2000Year 设置的时间，结果是： 12h0m0s

	// 时区
	// America/New_York
	// Asia/Shanghai
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		log.Panicln("转换时区失败")
	}
	fmt.Println(time.Now().In(loc)) // 这里需要配合t.In()

	// test time.time
	var tt time.Time
	var tw time.Weekday

	var ttp *time.Time
	var twp *time.Weekday

	p(tt)
	p(tw)
	p(ttp)
	p(twp)

}
