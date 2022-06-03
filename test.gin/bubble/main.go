package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-gorm/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)

var (
	DB *gorm.DB
)

// Tode Model
type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func initMySQL() (err error) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/bubble?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dsn)
	// 测试连通性
	err = DB.DB().Ping()
	return
	// return DB.DB().Ping()
}

func main() {
	// 创建数据库
	// sql: CREATE DATABASE bubble;
	// 连接数据库
	err := initMySQL()
	if err != nil {
		// 可以搞个日志
		panic(err)
	}
	// 模型绑定
	DB.AutoMigrate(&Todo{}) // todos
	defer DB.Close()        // 程序退出，关闭数据库
	// 遇事不决写注释
	r := gin.Default()
	// 告诉gin框架模板文件引用的静态文件去哪里找
	r.Static("/static", "static")
	// 告诉gin框架去哪里找模板文件
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// v1 api
	v1Group := r.Group("v1")
	{
		// 待办事项
		// 添加
		v1Group.POST("/todo", func(c *gin.Context) {
			// 前端页面填写待办事项，点击提交，会发请求到这里
			// 1. 从请求中把数据拿出来
			var todo Todo
			c.BindJSON(&todo)
			// 2. 存入数据库
			err := DB.Create(&todo).Error
			if err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, todo)
				// c.JSON(http.StatusOK, gin.H{
				// "code": 200,
				// "msg" : "success",
				// "data": todo,
				// })
			}
			// 3. 返回响应

		})
		// 查看所有的代办事项
		v1Group.GET("/todo", func(c *gin.Context) {
			// 查询todo这个表里的所有数据
			var todeList []Todo
			err := DB.Find(&todeList).Error
			if err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, todeList)
			}
		})
		// 查看某一个代办事项
		v1Group.GET("/todo/:id", func(c *gin.Context) {

		})
		// 修改某一个待办事项
		v1Group.PUT("/todo/:id", func(c *gin.Context) {
			id, ok := c.Params.Get("id")
			if !ok {
				c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
				return
			}
			var todo Todo
			if err = DB.Where("id=?", id).First(&todo).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
				return
			}
			c.BindJSON(&todo)
			if err = DB.Save(&todo).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, todo)
			}
		})
		// 删除
		v1Group.DELETE("/todo/:id", func(c *gin.Context) {
			id, ok := c.Params.Get("id")
			if !ok {
				c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
				return
			}
			if err = DB.Where("id=?", id).Delete(Todo{}).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, gin.H{id: "deleted"})
			}
		})
	}
	r.Run(":9090")
}
