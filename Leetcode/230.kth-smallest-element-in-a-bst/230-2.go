package _30_kth_smallest_element_in_a_bst

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func kthSmallest(root *TreeNode, k int) int {
	var res, rank int = 0, 0
	traverse(root, k, &res, &rank)
	return res
}

func traverse(root *TreeNode, k int, res, rank *int) {
	if root == nil {
		return
	}
	traverse(root.Left, k, res, rank)
	*rank++
	if k == *rank {
		*res = root.Val
		return
	}
	traverse(root.Right, k, res, rank)
}
