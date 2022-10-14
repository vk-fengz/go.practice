package _30_kth_smallest_element_in_a_bst

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res, rank int = 0, 0

func kthSmallest(root *TreeNode, k int) int {
	traverse(root, k)
	return res
}

func traverse(root *TreeNode, k int) {
	if root == nil {
		return
	}
	traverse(root.Left, k)
	rank++
	if k == rank {
		res = root.Val
		return
	}
	traverse(root.Right, k)
}
