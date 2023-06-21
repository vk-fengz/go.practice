package main

import (
	"fmt"
	"time"
)

func main() {
	var layout string = "2006-01-02 15:04:05"
	var timeStr string = "2019-12-12 15:22:12"
	var timeOnly string = "15:22:12"

	timeObj1, _ := time.Parse(layout, timeStr)
	fmt.Println(timeObj1)
	timeObj2, _ := time.ParseInLocation(layout, timeStr, time.Local)
	fmt.Println(timeObj2)

	tOnly, _ := time.Parse(time.TimeOnly, timeOnly)
	fmt.Println(tOnly)
	// use "15:22:12" construct new time days+2
	tNowCons := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day()+2, tOnly.Hour(), tOnly.Minute(), tOnly.Second(), 0, time.Now().Location())
	fmt.Println(tNowCons)
}
