package _54_russian_doll_envelopes

import (
	"fmt"
	"sort"
)

func maxEnvelopes(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		a, b := envelopes[i], envelopes[j]
		return a[0] < b[0] || a[0] == b[0] && a[1] > b[1]
	})

	f := make([]int, len(envelopes))
	for i, _ := range envelopes {
		f[i] = envelopes[i][1]
	}

	return lengthOfLIS(f)
}

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
