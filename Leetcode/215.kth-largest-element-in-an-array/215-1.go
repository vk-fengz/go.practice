package _15_kth_largest_element_in_an_array

func findKthLargest(nums []int, k int) int {
	shuffle(nums)
	lo, hi := 0, len(nums)-1
	// 升序排列, 转换成  排名第k 的元素
	k = len(nums) - k
	for lo <= hi {
		p := partition(nums, lo, hi)
		if p < k {
			lo = p + 1
		} else if p > k {
			hi = p - 1
		} else {
			return nums[p]
		}
	}
	return -1
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
