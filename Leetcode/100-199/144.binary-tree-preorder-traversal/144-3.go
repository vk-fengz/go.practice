package _44_binary_tree_preorder_traversal

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	res := preOrder(root)
	return res
}

// 迭代
func preOrder(root *TreeNode) (vals []int) {
	stack := []*TreeNode{}
	node := root
	for len(stack) > 0 || node != nil {
		for node != nil {
			vals = append(vals, node.Val)
			stack = append(stack, node)
			node = node.Left
		}

		node = stack[len(stack)-1].Right
		stack = stack[:len(stack)-1]
	}
	return
}
