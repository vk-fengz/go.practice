package _10_splitArray

// 与 1011 一致; 二分查找左侧最值
func splitArray(nums []int, m int) int {
	return shipWithinDays(nums, m)
}

// 返回能在 days 天内将传送带上的所有包裹送达的船的最低运载能力。
// 返回最小运载能力 -- 范围内查找
func shipWithinDays(weights []int, days int) int {
	// right 是开区间
	left, right := 0, 1
	// 确定 left right 边界
	for _, val := range weights {
		left = max(left, val)
		right += val
	}

	for left < right {
		mid := left + (right-left)>>1
		if fx(weights, mid) <= days {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

// 定义: 船的承载运力x, fx 天完成货物运输;  x--days
// fx 随着x的递增,而单调递减
func fx(weights []int, x int) int {
	// 返回days
	days := 0
	for i := 0; i < len(weights); {
		cap := x
		for i < len(weights) {
			if cap < weights[i] {
				break
			} else {
				cap -= weights[i]
			}
			i++
		}
		days++
	}
	return days
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
