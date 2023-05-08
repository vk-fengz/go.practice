package main

import "fmt"

func main() {
	books := make(map[string]string)
	books["a"] = "aaa"
	books["b"] = "bbb"
	fmt.Println(books)

	// 复制陷阱, 通过for range 复制到新map中
	aBook := books
	aBook["c"] = "ccc"
	fmt.Println(aBook)
	fmt.Println(books)
}
