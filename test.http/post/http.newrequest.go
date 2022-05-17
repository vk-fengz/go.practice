package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"unsafe"
)

type JsonPostSample struct {
}

func (this *JsonPostSample) SamplePost() {
	info := make(map[string]string)
	info["username"] = "******"
	info["password"] = "******"

	bytesData, err := json.Marshal(info)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(bytesData)

	reader := bytes.NewReader(bytesData)

	url := "http://xxxxxxxx.com" // 要访问的Url地址
	request, err := http.NewRequest("POST", url, reader)
	defer request.Body.Close() // 程序在使用完回复后必须关闭回复的主体
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	// 必须设定该参数,POST参数才能正常提交，意思是以json串提交数据

	client := http.Client{}
	resp, err := client.Do(request) // Do 方法发送请求，返回 HTTP 回复
	if err != nil {
		fmt.Println("22222", err.Error())
		return
	}

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("33333", err.Error())
		return
	}

	// byte数组直接转成string，优化内存
	str := (*string)(unsafe.Pointer(&respBytes))
	fmt.Println("44444", *str)

	// fmt.Println(string(respBytes))
}

func main() {
	a := new(JsonPostSample)
	a.SamplePost()

}
