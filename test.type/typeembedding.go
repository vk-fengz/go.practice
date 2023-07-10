// 实现 当内部类型和外部类型要
// 实现同一个接口  ==== 实现同一个接口
package main

import (
	"fmt"
)

// notifier是一个定义了
// 通知类行为的接口  == 接口定义的是行为;
type notifier interface {
	notify()
}

// 定义user用户类型
type user struct {
	name  string
	email string
}

// 通过user类型值的指针 调用方法
func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n",
		u.name,
		u.email)
}

// admin代表一个拥有权限的管理员用户
type admin struct {
	user
	level string
}

// 通过admin类型值指针调用的方法
func (a *admin) notify() {
	fmt.Printf("Sending admin email to %s<%s>\n",
		a.name,
		a.email)
}

// main()
func main() {
	ad := admin{
		user: user{
			name:  "john smith",
			email: "john@yahoo.com",
		},
		level: "super",
	}

	// 给admin用户发送一个通知
	// 接口的嵌入的内部类型实现并没有提升到外部类型
	sendNotification(&ad)

	// 直接访问内部类型的方法
	ad.user.notify()

	// 内部类型的方法没有提升
	ad.notify()
}

// 接受一个实现了notifier接口的值
// 并发送通知
func sendNotification(n notifier) {
	n.notify()
}

// 输出
/*
Sending admin email to john smith<john@yahoo.com>
Sending user email to john smith<john@yahoo.com>
Sending admin email to john smith<john@yahoo.com>
*/
