package _00_lengthOfLIS

// 动态规划
func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))

	// base case: dp数组初始化为1
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
	}

	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}
	res := 0
	for i := 0; i < len(nums); i++ {
		res = max(res, dp[i])
	}
	return res
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
