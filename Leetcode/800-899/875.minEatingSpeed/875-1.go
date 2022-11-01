package _75_minEatingSpeed

// 返回她可以在 h 小时内吃掉所有香蕉的最小速度 k（k 为整数）
func minEatingSpeed(piles []int, h int) int {
	// 香蕉的堆数目
	left, right := 1, 1000000000+1

	// 搜索左边界
	for left < right {
		mid := left + (right-left)/2
		if fx(piles, mid) <= h {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

// 吃的速度为x, 返回时间
func fx(piles []int, x int) int {
	hours := 0
	for i := 0; i < len(piles); i++ {
		hours += piles[i] / x
		if piles[i]%x > 0 {
			hours++
		}
	}
	return hours
}
