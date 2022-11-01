package _26_invert_binary_tree

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// invertTree 翻转二叉树
//	分解问题: 通过子问题(其中一个节点)推导出原问题的答案;
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	// 当前的X节点
	left := invertTree(root.Left)
	right := invertTree(root.Right)

	// 当前root节点
	root.Right = left
	root.Left = right

	return root
}
