package _22_coin_change

// 动态规划

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	// 填充dp table  amount+1
	dp[0] = 0
	for i := 1; i < amount+1; i++ {
		dp[i] = amount + 1
	}
	// 遍历所有可能
	for i := 1; i < amount+1; i++ {
		for j := 0; j < len(coins); j++ {
			if coins[j] <= i {
				dp[i] = min(dp[i], dp[i-coins[j]]+1)
			}
		}
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
