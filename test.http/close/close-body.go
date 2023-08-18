// https://mp.weixin.qq.com/s/OT5Jst4RNZrQ15a4TNy9OA
// 36、关闭HTTP的Response.Body
// 使用defer语句关闭资源时要注意nil值，在defer语句之前要进行nill值处理

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("https://api.ipify.org?format=json")

	if resp != nil {
		defer resp.Body.Close() // ok，即使不读取Body中的数据，即使Body是空的，也要调用close方法
	}

	//defer resp.Body.Close()       // （1）Error:在nil值判断之前使用，resp为nil时defer中的语句执行会引发空引用的panic
	if err != nil {
		fmt.Println(err)
		return
	}

	//defer resp.Body.Close()       // （2）Error:排除了nil隐患，但是出现重定向错误时，仍然需要调用close
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(body))
}
