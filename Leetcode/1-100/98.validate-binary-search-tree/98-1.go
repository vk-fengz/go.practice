package _8_validate_binary_search_tree

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return isBST(root, nil, nil)
}

// root节点   左子树的max < root.Val < 右子树min
// # 大小值 仅仅是 比较大小, 不涉及概念的东西; 仅仅是逻辑相关的
func isBST(root, min, max *TreeNode) bool {
	if root == nil {
		return true
	}

	if min != nil && root.Val <= min.Val {
		return false
	}
	if max != nil && root.Val >= max.Val {
		return false
	}

	return isBST(root.Left, min, root) && isBST(root.Right, root, max)
}
