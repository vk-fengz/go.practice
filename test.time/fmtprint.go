// go时间/时间戳操作大全: https://segmentfault.com/a/1190000019694913
package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("time.Now()：", time.Now())
	fmt.Println("2006-01-02 15:04:05：", time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println("format-999：", time.Now().Format("2006-01-02 15:04:05.999"))
	time.ParseDuration()

}
