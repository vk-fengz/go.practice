package _89_construct_binary_tree_from_preorder_and_postorder_traversal

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	if len(postorder) == 1 {
		return &TreeNode{Val: preorder[0]}
	}
	valueToIndex := make(map[int]int, len(postorder))
	for i, value := range postorder {
		valueToIndex[value] = i
	}
	index := valueToIndex[preorder[1]]

	root := &TreeNode{Val: preorder[0]}
	root.Left = constructFromPrePost(preorder[1:index+2], postorder[:index+1])
	root.Right = constructFromPrePost(preorder[index+2:], postorder[index+1:len(postorder)-1])
	return root
}
