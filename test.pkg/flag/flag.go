// go语言编程之旅 1.1
package main

import (
	"flag"
	"log"
)

func main() {
	var name string
	flag.StringVar(&name, "name", "Go 语言编程之旅", "帮助信息")
	flag.StringVar(&name, "n", "Go 语言编程之旅", "帮助信息")
	flag.Parse()

	log.Printf("name: %s", name)
}

// go run flag.go -name=eddycjy -n=煎鱼
// 2023/09/14 10:47:10 name: 煎鱼

// ==== 解析
// -flag：仅支持布尔类型。
// -flag x ：仅支持非布尔类型。
// -flag=x：均支持
