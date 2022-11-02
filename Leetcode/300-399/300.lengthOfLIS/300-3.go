package _00_lengthOfLIS

import "fmt"

// 动态规划
func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = 1
	res := 1
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j])
			}

		}
		dp[i] = dp[i] + 1
		res = max(res, dp[i])
	}
	fmt.Println(dp)
	return res
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
