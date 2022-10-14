package _6_unique_binary_search_trees

// Definition for a binary tree node.

// leetcode-book 中的解法
// # 使用嵌套的for循环来完成

// labuladong 解法
// 二维
func numTrees(n int) int {
	var memo [n + 1][n + 1]int
	return contBST(1, n, memo)
}

func contBST(lo, hi int, memo [][]int) int {
	if lo > hi {
		return 1
	}
	if memo[lo][hi] != 0 {
		return 1
	}

	res := 0
	for i := lo; i <= hi; i++ {
		left := contBST(lo, i-1, memo)
		right := contBST(i+1, hi, memo)
		res += left * right
	}
	memo[lo][hi] = res

	return res
}
