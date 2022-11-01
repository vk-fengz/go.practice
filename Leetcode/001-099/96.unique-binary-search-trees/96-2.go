package _6_unique_binary_search_trees

import "fmt"

func numTrees(n int) int {
	memo := make(map[string]int)
	return contBST(1, n, memo)
}

func contBST(lo, hi int, memo map[string]int) int {
	if lo > hi {
		return 1
	}
	lohi := fmt.Sprintf("%d-%d", lo, hi)
	if memo[lohi] != 0 {
		return memo[lohi]
	}

	res := 0
	for i := lo; i <= hi; i++ {
		left := contBST(lo, i-1, memo)
		right := contBST(i+1, hi, memo)
		res += left * right
	}
	memo[lohi] = res

	return res
}
