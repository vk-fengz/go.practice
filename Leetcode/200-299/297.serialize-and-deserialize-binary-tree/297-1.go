package _97_serialize_and_deserialize_binary_tree

import (
	"fmt"
	"strconv"
	"strings"
)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
	builder strings.Builder
	input   []string
}

func Constructor() Codec {
	return Codec{}
}

// 前序遍历
// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		this.builder.WriteString("#,")
		return ""
	}
	this.builder.WriteString(strconv.Itoa(root.Val) + ",")
	this.serialize(root.Left)
	this.serialize(root.Right)

	return this.builder.String()
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	fmt.Println(data)
	if len(data) == 0 {
		return nil
	}
	this.input = strings.Split(data, ",")
	return this.constructBinaryTree()
}

func (this *Codec) constructBinaryTree() *TreeNode {
	if this.input[0] == "#" {
		this.input = this.input[1:]
		return nil
	}
	val, _ := strconv.Atoi(this.input[0])
	this.input = this.input[1:]
	return &TreeNode{
		Val:   val,
		Left:  this.constructBinaryTree(),
		Right: this.constructBinaryTree(),
	}
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
