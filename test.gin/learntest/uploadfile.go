// https://geektutu.com/post/quick-go-gin.html
package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {

	// gin.Default()生成了一个实例，这个实例即 WSGI 应用程序
	r := gin.Default()

	// ===============================

	// 1. 上传 单个文件
	r.POST("/upload1", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		// c.SaveUploadedFile(file, dst)
		c.String(http.StatusOK, "%s uploaded!", file.Filename)
	})

	// 2. 上传多个文件
	r.POST("/upload2", func(c *gin.Context) {
		// Multipart form
		form, _ := c.MultipartForm()
		files := form.File["upload[]"]

		for _, file := range files {
			log.Println(file.Filename)
			// c.SaveUploadedFile(file, dst)
		}
		c.String(http.StatusOK, "%d files uploaded!", len(files))
	})

	// ================================ run
	r.Run("localhost:9999")
}
