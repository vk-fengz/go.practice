package main

import (
	"fmt"
	"time"
)

func main() {
	var layout string = "2006-01-02 15:04:05"
	var timeStr string = "2019-12-12 15:22:12"
	var timeOnly string = "15:22:12"
	var R1123Z string = "Mon, 02 Jan 2006 15:04:05 +0800"

	timeObj1, _ := time.Parse(layout, timeStr)
	fmt.Println(timeObj1)
	timeObj2, _ := time.ParseInLocation(layout, timeStr, time.Local)
	fmt.Println(timeObj2)

	tOnly, _ := time.Parse(time.TimeOnly, timeOnly)
	fmt.Println(tOnly)
	fmt.Println("tOnly after: ", tOnly.After(time.Now()))

	// use "15:22:12" construct new time days+2
	tNowCons := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), tOnly.Hour(), tOnly.Minute(), tOnly.Second(), 0, time.Now().Location())
	fmt.Println("构建时间after now: ", tNowCons.After(time.Now()))
	fmt.Println("构建时间", tNowCons)
	fmt.Println(&tNowCons) // 2023-06-21 15:22:12 +0800 CST

	// 取地址失败
	ptnowcons := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day()+2, tOnly.Hour(), tOnly.Minute(), tOnly.Second(), 0, time.Now().Location())
	fmt.Println("value nowCons", ptnowcons)
	fmt.Println("value nowCons", &ptnowcons)

	//RFC1123Z
	fmt.Println("--------- RFC1123Z")
	rfz, _ := time.Parse(time.RFC1123Z, R1123Z)
	fmt.Println("RFC1123Z dateTime: ", rfz.Month())
}

// 2019-12-12 15:22:12 +0000 UTC
// 2019-12-12 15:22:12 +0800 CST
// 0000-01-01 15:22:12 +0000 UTC
// tOnly after:  false
// 构建时间after now:  false
// 构建时间 2023-07-17 15:22:12 +0800 CST
// 2023-07-17 15:22:12 +0800 CST
// value nowCons 2023-07-19 15:22:12 +0800 CST
// value nowCons 2023-07-19 15:22:12 +0800 CST
// --------- RFC1123Z
// RFC1123Z dateTime:  January
