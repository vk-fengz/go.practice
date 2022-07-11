package _44_binary_tree_preorder_traversal

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func preorderTraversal(root *TreeNode) []int {
	res := []int{}
	preOrder(root, &res)
	return res
}

func preOrder(root *TreeNode, output *[]int) {
	if root != nil {
		*output = append(*output, root.Val)
		preOrder(root.Left, output)
		preOrder(root.Right, output)
	}
}
