// https://mp.weixin.qq.com/s/OT5Jst4RNZrQ15a4TNy9OA
// （3）全局关闭http连接重用。

// 对于相同http server发送多个请求时，适合保持网络连接；
// 但对于短时间内向多个HTTP服务器发送一个或两个请求时，最好在每次接收到服务端响应后关闭网络链接
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// 全局关闭http连接重用
	tr := &http.Transport{DisableKeepAlives: true}
	client := &http.Client{Transport: tr}

	resp, err := client.Get("http://golang.org")
	if resp != nil {
		defer resp.Body.Close()
	}

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(resp.StatusCode)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(len(string(body)))
}
