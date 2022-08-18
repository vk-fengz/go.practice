package _12_sort_an_array

import "math/rand"

func sortArray(nums []int) []int {
	return quickSort(nums)
}

func quickSort(nums []int) []int {
	// 为了避免出现耗时的极端情况，先随机打乱
	shuffle(nums)
	// 排序整个数组（原地修改）
	sort(nums, 0, len(nums)-1)
	return nums
}

func sort(nums []int, lo, hi int) []int {
	if lo >= hi {
		return nil
	}

	p := partition(nums, lo, hi)
	sort(nums, lo, p-1)
	sort(nums, p+1, hi)

	return nums[lo:hi]
}

func partition(nums []int, lo, hi int) int {
	pivot := nums[lo]
	// 关于区间的边界控制需格外小心，稍有不慎就会出错
	// 我这里把 i, j 定义为开区间，同时定义：
	// [lo, i) <= pivot；(j, hi] > pivot
	// 之后都要正确维护这个边界区间的定义
	i := lo + 1
	j := hi
	// 当 i > j 时结束循环，以保证区间 [lo, hi] 都被覆盖
	for i <= j {
		for i < hi && nums[i] <= pivot {
			i++
		}
		for j > lo && nums[j] > pivot {
			j--
		}
		if i >= j {
			break
		}
		swap(nums, i, j)
	}

	// 将 pivot 放到合适的位置，即 pivot 左边元素较小，右边元素较大
	swap(nums, lo, j)
	return j
}

func shuffle(nums []int) {
	n := len(nums)
	for i, _ := range nums {
		r := i + rand.Intn(n-i)
		swap(nums, i, r)
	}
}

func swap(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}
