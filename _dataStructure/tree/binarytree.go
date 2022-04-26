package main

import (
	"fmt"
)

// 二叉树
type TreeNode struct {
	Data  string    // 节点用来存放数据
	Left  *TreeNode // 左子树
	Right *TreeNode // 右字树
}

// 先序遍历
func PreOrder(tree *TreeNode) {
	if tree == nil {
		return
	}

	// 先打印根节点
	fmt.Print(tree.Data, " ")
	// 再打印左子树
	PreOrder(tree.Left)
	// 再打印右字树
	PreOrder(tree.Right)
}

// 中序遍历
func MidOrder(tree *TreeNode) {
	if tree == nil {
		return
	}

	MidOrder(tree.Left)
	fmt.Print(tree.Data, " ")
	MidOrder(tree.Right)
}

// 后序遍历
func PostOrder(tree *TreeNode) {
	if tree == nil {
		return
	}

	MidOrder(tree.Left)
	MidOrder(tree.Right)
	fmt.Print(tree.Data, " ")
}

func main() {
	t := &TreeNode{Data: "A"}
	t.Left = &TreeNode{Data: "B"}
	t.Right = &TreeNode{Data: "C"}
	t.Left.Left = &TreeNode{Data: "D"}
	t.Left.Right = &TreeNode{Data: "E"}
	t.Right.Left = &TreeNode{Data: "F"}

	fmt.Println("先序排序：")
	PreOrder(t)
	fmt.Println("\n中序排序：")
	MidOrder(t)
	fmt.Println("\n后序排序")
	PostOrder(t)
}
