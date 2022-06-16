package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Class struct {
	gorm.Model
	ClassName string
	Students  []Student // 1个班级有多个学生
}

type Student struct {
	gorm.Model
	StudentName string
	ClassID     uint   // 1个学生属于多个班级
	IDCard      IDCard // 1个学生拥有1张身份证
	// 多对多, 通过中间表关联
	Teachers []Teacher `gorm:"many2many:student_teachers;"`
}

type IDCard struct {
	gorm.Model
	Num       int
	StudentID uint // 1张身份证属于1个学生
}
type Teacher struct {
	gorm.Model
	TeacherName string
	// 多对多, 通过中间表关联
	Students []Student `gorm:"many2many:student_teachers;"`
}

func main() {
	CreateTable()
}

func CreateTable() {
	dsn := "root:lzy123456@tcp(127.0.0.1:3306)/ginclass?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db.AutoMigrate(&Teacher{}, &Class{}, &Student{}, &IDCard{})

	i := IDCard{
		Num: 123456,
	}
	t := Teacher{
		TeacherName: "老师傅",
	}
	s := Student{
		StudentName: "qm",
		IDCard:      i,
		Teachers:    []Teacher{t},
	}
	c := Class{
		ClassName: "奇妙的班级",
		Students:  []Student{s},
	}

	db.Create(&c)
	db.Create(&t)
}
