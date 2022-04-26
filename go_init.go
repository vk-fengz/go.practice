package main

import "fmt"

var a = hello()
func hello() int {
	fmt.Println("hello")
	return 0
}
func init() {
	fmt.Println("world")
}

func main() {
	fmt.Println("Hello World")
	fmt.Println(a)   //return 0
}

