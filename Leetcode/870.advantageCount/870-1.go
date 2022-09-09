package _70_advantageCount

import (
	"fmt"
	"sort"
)

// 升序排列nums1 降序排列nums2
func advantageCount(nums1 []int, nums2 []int) []int {

	n := len(nums1)
	sort.Ints(nums1)
	sortedB := make([]int, n)
	for i := range sortedB {
		sortedB[i] = i
	}
	// B 排序
	sort.Slice(sortedB, func(i, j int) bool {
		return nums2[sortedB[i]] < nums2[sortedB[j]]
	})
	// 打印输出
	fmt.Println(sortedB)

	useless, i, res := make([]int, 0), 0, make([]int, n)
	for _, index := range sortedB {
		b := nums2[index]
		for i < n && nums1[i] <= b {
			useless = append(useless, nums1[i])
			i++
		}
		if i < n {
			res[index] = nums1[i]
			i++
		} else {
			res[index] = useless[0]
			useless = useless[1:]
		}
	}
	return res
}
