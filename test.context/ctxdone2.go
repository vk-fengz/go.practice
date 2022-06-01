// done()是监听cancel或超时操作。

package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	type Output struct {
		Num int
	}

	ms := make(chan context.Context)

	go func() {
		for i := 0; i < 100; i++ {
			time.Sleep(time.Duration(5) * time.Second)

			// 添加context内容
			u := &Output{
				Num: i,
			}
			c := context.WithValue(context.Background(), "user", u)

			// 取消实现context结束
			c, cl := context.WithCancel(c)
			cl()

			ms <- c
		}
	}()

	for m := range ms {
		select {
		case <-m.Done():
			fmt.Println(fmt.Sprintf("结束原因：%s, 输出结果：%d", m.Err(), m.Value("user").(*Output).Num))
		default:
		}
	}
}
