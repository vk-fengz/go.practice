package main

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/upload", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		name := c.PostForm("name") // 获取其他数据

		// 将文件写到本地
		in, _ := file.Open()
		defer in.Close()
		out, _ := os.Create("./" + file.Filename)
		defer out.Close()
		io.Copy(out, in)

		c.JSON(200, gin.H{
			"file": file,
			"name": name,
		})
	})
	router.Run(":8080")
}
