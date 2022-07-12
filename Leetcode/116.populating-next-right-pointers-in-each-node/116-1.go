package _16_populating_next_right_pointers_in_each_node

// Definition for a Node.
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	traverse(root.Left, root.Right)
	return root
}

func traverse(left, right *Node) {
	if left == nil || right == nil {
		return
	}
	// 前序位置, 连接两个节点
	left.Next = right

	// 遍历框架
	traverse(left.Left, left.Right)
	traverse(right.Left, right.Right)
	// 跨父节点 连接
	traverse(left.Right, right.Left)

}
