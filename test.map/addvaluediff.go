package main

import "fmt"

func main() {
	books := make(map[string]string, 2)
	books["a"] = "aaa"
	books["b"] = "bbb"
	fmt.Println("The original map:")
	fmt.Println(books)
	books["dd"] = "dddd"
	fmt.Println("add value after: ", books) // 突破len限制
	fmt.Println("--------")

	// 复制陷阱, 通过for range 复制到新map中
	aBook := books
	aBook["c"] = "ccc"
	fmt.Println("copy map aBook: ", aBook)
	fmt.Println("origin map: ", books)
}
