// 直接插入

package main

import "fmt"

func insertSort(list []int) {
	n := len(list)
	// 进行 N-1 轮迭代
	for i := 1; i <= n-1; i++ {
		deal := list[i] // 待排序的数
		j := i - 1      // 待排序的数左边的第一个数的位置

		// 如果第一次比较, 比左边的已排好序的第一个数小, 那么进入处理
		if deal < list[j] {
			// 一直往左边找, 比待排序大的数都往后挪, 腾空位给待排序插入
			for ; j >= 0 && deal < list[j]; j-- {
				list[j+1] = list[j] // 某数后移, 给待排序留空位
			}
			list[j+1] = deal // 结束了, 待排序的数插入空位
		}
	}
}

func main() {
	list := []int{5}
	insertSort(list)
	fmt.Println(list)

	list1 := []int{5, 9}
	insertSort(list1)
	fmt.Println(list1)

	list11 := []int{5, 9}
	insertS(list11)
	fmt.Println(list11)

	list2 := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
	insertSort(list2)
	fmt.Println(list2)

	// insertS
	list22 := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
	insertS(list22)
	fmt.Println(list22)
}

// 手撕 insertSort 扑克牌 快速插入排序
func insertS(array []int) {
	n := len(array)
	for i := 1; i < n; i++ {
		j := i - 1
		temp := array[i]
		if temp < array[j] {
			for ; j >= 0 && temp < array[j]; j-- {
				array[j+1] = array[j]
			}
			array[j+1] = temp
		}
	}
}
