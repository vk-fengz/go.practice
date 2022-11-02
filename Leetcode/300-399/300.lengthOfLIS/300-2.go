package _00_lengthOfLIS

// 纸牌游戏方法
func lengthOfLIS(nums []int) int {
	// 记录每个牌堆上面的数字
	top := make([]int, len(nums))
	// piles 牌堆数
	piles := 0

	for i := 0; i < len(nums); i++ {
		poker := nums[i]
		// 搜索左边界的二分查找
		left, right := 0, piles
		for left < right {
			mid := left + (right-left)>>1
			if top[mid] > poker {
				right = mid
			} else if top[mid] < poker {
				left = mid + 1
			} else {
				right = mid
			}
		}
		if left == piles {
			piles += 1
		}
		top[left] = poker
	}
	return piles
}
