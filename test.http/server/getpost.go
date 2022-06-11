package main

import (
	"fmt"
	"log"
	"net/http"
)

func sayHelloHandler(w http.ResponseWriter, r *http.Request) {

	r.ParseForm() // 解析url传递的参数，对于POST则解析响应包的主体（request body）
	// 注意:如果没有调用ParseForm方法，下面无法获取表单的数据
	uid := r.Form["uid"]
	fmt.Println(uid)
}

func main() {
	http.HandleFunc("/", sayHelloHandler) //   设置访问路由
	log.Fatal(http.ListenAndServe(":8081", nil))
}
