package main

import (
	"fmt"
	"net/http"
)
// 首先注册路由，然后创建服务并开启监听即可。
type indexHandler struct {
	content string
}

func (ih *indexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, ih.content)
}

func main() {
	http.Handle("/", &indexHandler{content: "hello world!"})    // 注册路由router
	http.ListenAndServe(":8001", nil)
}
