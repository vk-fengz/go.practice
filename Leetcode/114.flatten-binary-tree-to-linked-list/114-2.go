package _14_flatten_binary_tree_to_linked_list

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 分解问题 的思维模式, 单就一个点 来实现算法;
func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	// 把节点的左右子树拉平
	flatten(root.Left)
	flatten(root.Right)

	// 后序遍历位置,
	left, right := root.Left, root.Right
	root.Left = nil3
	root.Right = left

	// 原右子树拼接到当前右子树的末端
	p := root
	for p.Right != nil {
		p = p.Right
	}
	p.Right = right

}
