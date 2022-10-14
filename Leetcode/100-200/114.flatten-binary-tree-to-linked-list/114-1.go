package _14_flatten_binary_tree_to_linked_list

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 二叉树前序遍历, 一遍出结果,
// 占用内存大, -2 原地实现

func flatten(root *TreeNode) {
	list := preorderTraversal(root)
	for i := 0; i < len(list)-1; i++ {
		cur, curr := list[i], list[i+1]
		cur.Left, cur.Right = nil, curr
	}
}

func preorderTraversal(node *TreeNode) []*TreeNode {
	var list []*TreeNode
	if node != nil {
		list = append(list, node)
		list = append(list, preorderTraversal(node.Left)...)
		list = append(list, preorderTraversal(node.Right)...)
	}
	return list
}
