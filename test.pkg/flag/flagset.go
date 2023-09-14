// refer: https://golang2.eddycjy.com/posts/ch1/01-simple-flag/
package main

import (
	"flag"
	"log"
)

var name string

func main() {
	flag.Parse() //  flag.Parse 方法，将命令行解析为定义的标志，便于我们后续的参数使用。

	args := flag.Args()
	if len(args) <= 0 {
		return
	}

	switch args[0] {
	case "go":
		goCmd := flag.NewFlagSet("go", flag.ExitOnError) // parse 方法的异常分流处理; 可处理复解析、异常处理等，均直接由该方法处理
		goCmd.StringVar(&name, "name", "Go 语言", "帮助信息")
		_ = goCmd.Parse(args[1:]) // flagset.parse
	case "php":
		phpCmd := flag.NewFlagSet("php", flag.ExitOnError)
		phpCmd.StringVar(&name, "n", "PHP 语言", "帮助信息")
		_ = phpCmd.Parse(args[1:])
	}

	log.Printf("name: %s", name)
}
