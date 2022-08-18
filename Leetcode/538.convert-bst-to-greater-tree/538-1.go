package _38_convert_bst_to_greater_tree

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// BST二叉搜索树

func convertBST(root *TreeNode) *TreeNode {
	sum := 0
	traverse(root, &sum)
	return root
}

func traverse(root *TreeNode, sum *int) {
	if root != nil {
		traverse(root.Right, sum)
		*sum += root.Val
		root.Val = *sum
		traverse(root.Left, sum)
	}
}
