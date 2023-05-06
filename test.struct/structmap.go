package main

import (
	"fmt"
)

type Config struct {
	num  int
	link map[string]string
}

func main() {
	config := Config{
		num:  1,
		link: fill(),
	}

	fmt.Println(config)
	if val, ok := config.link["a"]; ok {
		fmt.Println(val)
	} else {
		fmt.Println("not contains value a ")
	}

	// illegal:  nil map, ==> panic
	// config.link["a"] = "aaa"
	// fmt.Println(config.link)

	// legal
	config.link = map[string]string{
		"a": "aaa",
	}
	fmt.Println(config)
}

func fill() map[string]string {
	return map[string]string{
		"b": "bbb",
	}

	// test
}
