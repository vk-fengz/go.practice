package _05_construct_binary_tree_from_preorder_and_inorder_traversal

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	index, value := 0, 0
	for index, value = range inorder {
		if value == preorder[0] {
			break
		}
	}

	root := &TreeNode{Val: value}
	root.Left = buildTree(preorder[1:index+1], inorder[:index])
	root.Right = buildTree(preorder[index+1:], inorder[index+1:])
	return root
}

// 考虑使用hash map 建立元素索引
