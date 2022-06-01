// https://geektutu.com/post/quick-go-gin.html
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	// gin.Default()生成了一个实例，这个实例即 WSGI 应用程序
	r := gin.Default()
	// ===================================

	type student struct {
		Name string
		Age  int8
	}

	r.LoadHTMLGlob("templates/*")

	stu1 := &student{Name: "Geektutu", Age: 20}
	stu2 := &student{Name: "Jack", Age: 22}
	r.GET("/arr", func(c *gin.Context) {
		c.HTML(http.StatusOK, "arr.tmpl", gin.H{
			"title":  "Gin",
			"stuArr": [2]*student{stu1, stu2},
		})
	})

	// =================================== run
	r.Run("localhost:9999") // 为空时,  listen and serve on 0.0.0.0:8080
}
