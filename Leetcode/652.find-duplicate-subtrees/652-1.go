package _52_find_duplicate_subtrees

import "strconv"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 后续遍历实现

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	mem := make(map[string]int)
	res := make([]*TreeNode, 0)
	var traverse func(root *TreeNode) string
	traverse = func(root *TreeNode) string {
		if root == nil {
			return "#"
		}

		left := traverse(root.Left)
		right := traverse(root.Right)
		subtree := left + "," + right + "," + strconv.Itoa(root.Val)

		freq := mem[subtree]
		if freq == 1 {
			res = append(res, root)
		}
		mem[subtree] = freq + 1
		return subtree
	}

	traverse(root)
	return res
}
