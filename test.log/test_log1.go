// 这个示例程序展示如何使用最基本的 log 包
package main

import (
	"log"
)

func init() {
	log.SetPrefix("TRACE: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}

func main() {
	// Println 写到标准日志记录器
	log.Println("message")

	// Fatalln 在调用 Println()之后会接着调用 os.Exit(1)
	log.Fatalln("fatal message")

	// Panicln 在调用 Println()之后会接着调用 panic()
	log.Panicln("panic message")
}
