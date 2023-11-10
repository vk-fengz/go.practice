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

	p("----------- 构造时间: ")
	// 构造时间，Date 的最后一个参数是时区
	myTime := time.Date(1996, time.April, 27, 0, 0, 0, 0, time.UTC)
	myTimeLocal := time.Date(1996, time.April, 27, 0, 0, 0, 0, time.Now().Location())
	// 两天后
	twoDayAftermyTimeLocal := myTimeLocal.AddDate(0, 0, 2)
	p("我构造的时间：", myTime)
	p("我构造的时间 localtime：", myTimeLocal)
	p("我构造的时间 两天后 localtime：", twoDayAftermyTimeLocal)

	p("----------- 时间对比 加减:")
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
	p("----------- 时区 ")
	// America/New_York
	// Asia/Shanghai
	loc, err := time.LoadLocation("Asia/Shanghai") // windows系统会出错,依赖tzdata
	if err != nil {
		log.Panicln("转换时区失败")
	}
	fmt.Println("local 当前时间:", time.Now().In(loc)) // 这里需要配合t.In()

	tUTC := time.Unix(myTimeLocal.Unix()+10000, 0).UTC()
	tLoc := time.Unix(myTimeLocal.Unix()+10000, 0).Local()
	p("myTime add sec UTC:", tUTC)
	p("myTime add sec localTimezone:", tLoc)

	// method 2    timezone
	p("----------- 设置时区 方法2")
	var cstZone = time.FixedZone("CST", 8*3600) // 东八区, 设置时区推荐方法
	p("==> set CST zone")
	time.Local = cstZone
	tSH := time.Now().Local()
	p("tSH: ", tSH)
	p("tSH localtion: ", tSH.Location())

	time.Sleep(5 * time.Second)
	p("==> set UTC zone")
	time.Local = time.UTC
	tUTCz := time.Now().Local()
	p("tUTCz: ", tUTCz)
	p("tUTCz to CST: ", tUTCz.In(cstZone))
	p("tUTCz localtion: ", tUTCz.Location())
	// time compare
	p("==> time compare: tUTCz.After(tSH) == ", tUTCz.After(tSH))
	p("tUTCz=", tUTCz)
	p("tSH  =", tSH)

	p("-------- test var time")
	tt := time.Time{}
	var tw time.Weekday

	var ttp *time.Time
	var twp *time.Weekday

	p(tt)
	p(tw)
	p(ttp)
	p(twp)

}

// aint的地址: 0xc0000ac000 10
// 当前时间： 2023-11-03 17:54:53.670630828 +0800 CST m=+0.000137122
// 当前时间地址 2023-11-03 17:54:53.670630828 +0800 CST m=+0.000137122
// 当前时间utc： 2023-11-03 09:54:53.670786144 +0000 UTC
// 当前年份： 2023
// 当前是周几： Friday
// 当前是周几 pointer： Friday
// 明天是周几： Saturday
// ----------- 构造时间:
// 我构造的时间： 1996-04-27 00:00:00 +0000 UTC
// 我构造的时间 localtime： 1996-04-27 00:00:00 +0800 CST
// 我构造的时间 两天后 localtime： 1996-04-29 00:00:00 +0800 CST
// ----------- 时间对比 加减:
// 2000 年在 3000 年之前吗： true
// end - start 的时间，结果是： 12h0m0s
// ----------- 时区
// local 当前时间: 2023-11-03 17:54:53.670933546 +0800 CST
// myTime add sec UTC: 1996-04-26 18:46:40 +0000 UTC
// myTime add sec localTimezone: 1996-04-27 02:46:40 +0800 CST
// -------- test var time
// 0001-01-01 00:00:00 +0000 UTC
// Sunday
// <nil>
// <nil>
