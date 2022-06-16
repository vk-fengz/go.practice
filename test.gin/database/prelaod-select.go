package main

func main() {
	// 连接数据库
	dsn := "root:lzy123456@tcp(127.0.0.1:3306)/ginclass?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db.AutoMigrate(&Teacher{}, &Class{}, &Student{}, &IDCard{})

	r := gin.Default()

	// 新增学生
	r.POST("/student", func(c *gin.Context) {
		var student Student
		_ = c.BindJSON(&student)
		db.Create(&student)
	})

	// 根据id查询学生
	r.GET("/student/:id", func(c *gin.Context) {
		id := c.Param("id")
		var student Student
		_ = c.BindJSON(&student)
		// db.First(&student, "id = ?", id) // 无法加载出关联的数据
		db.Preload("Teachers").Preload("IDCard").
			First(&student, "id = ?", id) // 可以查出关联数据
		c.JSON(200, gin.H{"student": student})
	})

	// 根据id查询班级
	r.GET("/class/:id", func(c *gin.Context) {
		id := c.Param("id")
		var class Class
		// db.First(&class, "id = ?", id) // 无法加载出关联的数据
		db.Preload("Students").Preload("Students.Teachers").Preload("Students.IDCard").
			First(&class, "id = ?", id) // 可以查出关联数据
		c.JSON(200, gin.H{"class": class})
	})

	r.Run(":8888")
}
