// https://geektutu.com/post/quick-go-gin.html

package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	// gin.Default()生成了一个实例，这个实例即 WSGI 应用程序
	r := gin.Default()

	// ========================== 使用示例 ====================
	// 1. GET hello world
	// r.GET("/", func(c *gin.Context) {
	// 	c.String(200, "Hello, Geektutu")
	// })

	// 2. 无参数
	// curl http://localhost:9999/
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Who are you?")
	})

	// 3. 解析路径参数 动态路由
	// 匹配 /user/geektutu
	// curl http://localhost:9999/user/geektutu
	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	// 4. 获取 Query 参数
	// 匹配 users?name=xxx&role=xxx，role可选
	// curl "http://localhost:9999/users?name=Tom&role=student"
	r.GET("/users", func(c *gin.Context) {
		name := c.Query("name")
		role := c.DefaultQuery("role", "teacher")
		c.String(http.StatusOK, "%s is a %s", name, role)
	})

	// 5. 获取 post 参数
	// POST
	r.POST("/form", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.DefaultPostForm("password", "000000") // 可设置默认值

		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
		})
	})

	// 6. Query和POST混合参数
	// GET 和 POST 混合
	// curl "http://localhost:9999/posts?id=9876&page=7"  -X POST -d 'username=geektutu&password=1234'
	r.POST("/posts", func(c *gin.Context) {
		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		username := c.PostForm("username")
		password := c.DefaultPostForm("username", "000000") // 可设置默认值

		c.JSON(http.StatusOK, gin.H{
			"id":       id,
			"page":     page,
			"username": username,
			"password": password,
		})
	})

	// 7. Map参数(字典参数)
	// $ curl -g "http://localhost:9999/post?ids[Jack]=001&ids[Tom]=002" -X POST -d 'names[a]=Sam&names[b]=David'
	// {"ids":{"Jack":"001","Tom":"002"},"names":{"a":"Sam","b":"David"}}
	r.POST("/post", func(c *gin.Context) {
		ids := c.QueryMap("ids")
		names := c.PostFormMap("names")

		c.JSON(http.StatusOK, gin.H{
			"ids":   ids,
			"names": names,
		})
	})

	// 8. 重定向
	//  curl -i http://localhost:9999/redirect
	// curl "http://localhost:9999/goindex"
	r.GET("/redirect", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/index")
	})

	r.GET("/goindex", func(c *gin.Context) {
		c.Request.URL.Path = "/"
		r.HandleContext(c)
	})

	// 9. 分组路由
	// $ curl http://localhost:9999/v1/posts
	// {"path":"/v1/posts"}
	// $ curl http://localhost:9999/v2/posts
	// {"path":"/v2/posts"}

	// group routes 分组路由
	defaultHandler := func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"path": c.FullPath(),
		})
	}
	// group: v1
	v1 := r.Group("/v1")
	{
		v1.GET("/posts", defaultHandler)
		v1.GET("/series", defaultHandler)
	}
	// group: v2
	v2 := r.Group("/v2")
	{
		v2.GET("/posts", defaultHandler)
		v2.GET("/series", defaultHandler)
	}

	// ==============================================
	// == run
	r.Run("localhost:9999") // 为空时,  listen and serve on 0.0.0.0:8080
}
