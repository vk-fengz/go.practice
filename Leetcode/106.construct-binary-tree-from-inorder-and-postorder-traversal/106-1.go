package _06_construct_binary_tree_from_inorder_and_postorder_traversal

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	index, value := 0, 0
	for index, value = range inorder {
		if value == postorder[len(postorder)-1] {
			break
		}
	}

	root := &TreeNode{Val: value}
	root.Left = buildTree(inorder[:index], postorder[:index])
	root.Right = buildTree(inorder[index+1:], postorder[index:len(postorder)-1])
	return root
}
