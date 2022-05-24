// go 实现发送post请求的两种方法
// https://blog.csdn.net/YMY_mine/article/details/98496009
// 可以用 apifox 等直接生成

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"unsafe"
)

func main() {
	song := make(map[string]string)
	song["username"] = "******"
	song["password"] = "******"
	bytesData, _ := json.Marshal(song)

	res, err := http.Post("http://127.0.0.1:8080",
		"application/json;charset=utf-8", bytes.NewBuffer([]byte(bytesData)))
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
	}

	defer res.Body.Close()

	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
	}

	// fmt.Println(string(content))
	str := (*string)(unsafe.Pointer(&content)) // 转化为string,优化内存
	fmt.Println(*str)
}
