package _54_maximum_binary_tree

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	var dfs func(nums []int) *TreeNode
	dfs = func(nums []int) *TreeNode {
		if len(nums) == 0 {
			return nil
		}

		maxIndex, maxValue := 0, 0
		// 查找最大值及其索引位置
		for index, value := range nums {
			if value > maxValue {
				maxIndex = index
				maxValue = value
			}
		}
		// 构建左右子树
		root := &TreeNode{Val: maxValue}
		root.Left = dfs(nums[:maxIndex])
		root.Right = dfs(nums[maxIndex+1:])
		return root
	}
	return dfs(nums)
}
