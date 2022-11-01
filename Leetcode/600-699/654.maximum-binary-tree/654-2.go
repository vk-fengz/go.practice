package _54_maximum_binary_tree

import "math"

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
		// MaxIndex就是动态记录最大值的位置
		MaxIndex := 0
		// MaxIndex就是动态记录最大值是多少
		MaxIndex := math.MinInt64
		root := &TreeNode{}
		for i, v := range nums {
			if v > maxValue {
				maxValue = v
				MaxIndex = i
			}
		}
		// 上面的for循环就是为了求出根节点（nums中的最大值），和他的坐标
		root.Val = maxValue
		root.Left = dfs(nums[:MaxIndex])
		root.Right = dfs(nums[MaxIndex+1:])
		return root
	}
	return dfs(nums)
}

// 作者：gracious-vvilson1bb
// 链接：https://leetcode.cn/problems/maximum-binary-tree/solution/-by-gracious-vvilson1bb-sa5e/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
