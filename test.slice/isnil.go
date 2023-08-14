package main

import (
	"fmt"
)

func main() {
	var num []int

	if num == nil {
		fmt.Println("切片是空的, nil")
		fmt.Println("address:", &num)
	}

	/* 打印原始切片 */
	fmt.Println("numbers ==", num)

}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}

// 切片是空的, nil
// address: &[]
// numbers == []
