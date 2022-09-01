package main

import "fmt"

func main() {

	// 共有 m 个 [n]
	myArray := [3][4]int{{1, 2, 3, 4}, {1, 2, 3, 4}, {1, 2, 3, 4}}

	// 打印一维数组长度   m  高度
	fmt.Println(len(myArray))
	// 打印二维数组长度   n  长度
	fmt.Println(len(myArray[1]))

	mySlice := myArray[0]
	fmt.Println(mySlice)

	fmt.Println(len(mySlice))
}

// 输出
// 3
// 4
// [1 2 3 4]
// 4
