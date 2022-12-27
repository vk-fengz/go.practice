package main

import (
	"fmt"
)

// Unique
// 对于一个  有序数组   arr ，需要原地删除重复出现的元素，使每个元素只出现一次 ，
// 不要使用额外的数组空间，并在使用 O(1) 额外空间的条件下完成。
func Unique(arr []int) []int {
	left := 0
	for right := 1; right < len(arr); right++ {
		if arr[left] != arr[right] {
			left++
			arr[left] = arr[right]
		}
	}
	return arr[:left+1]
}

func main() {
	s := []int{1, 2, 2, 3} // 创建一个int类型切片并赋值
	res := Unique(s)
	fmt.Println(res) // [1 2 3]
}
