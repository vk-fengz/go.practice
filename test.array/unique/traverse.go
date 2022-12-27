package main

import (
	"fmt"
)

// Unique 从后向前进行遍历，拿最后一项自后往前逐个进行比较，当遇到有相同项时移除最后一项，同时跳出比较。这里因为自后往前比较，下标以及总长度的问题不用处理。这是数组去重最佳的写法，效率较高，留下的非重复项也是前面的项。
func Unique(arr []int) []int {
	arr_len := len(arr) - 1
	for ; arr_len > 0; arr_len-- {
		for i := arr_len - 1; i >= 0; i-- {
			if arr[arr_len] == arr[i] {
				arr = append(arr[:arr_len], arr[arr_len+1:]...)
				fmt.Println(arr)
				break
			}
		}
		fmt.Println(arr)
	}
	return arr
}
func main() {
	s := []int{1, 2, 4, 2, 1, 2, 3} // 创建一个int类型切片并赋值
	res := Unique(s)
	fmt.Println(res) // [1 2 3]
}
