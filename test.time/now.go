// test now time and now func
package main

import (
	"fmt"
	"time"
)

func main() {
	nowfunc := time.Now
	told := time.Now()
	toldfunc := nowfunc()

	fmt.Println("told    :", told)
	fmt.Println("toldfunc:", toldfunc)

	// wait 10s
	time.Sleep(10 * time.Second)
	tnew := time.Now()
	tnewfunc := nowfunc()
	fmt.Println("tnew    :", tnew)
	fmt.Println("tnewfunc:", tnewfunc) //和 now func 一致, 保持时间更新
	fmt.Println("told    :", told)     //固定时刻

}
