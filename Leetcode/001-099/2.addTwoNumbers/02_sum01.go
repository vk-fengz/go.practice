package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/*
// 示例:
输入：l1 = [2,4,3], l2 = [5,6,4]
输出：[7,0,8]
解释：342 + 465 = 807.
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1, l2 *ListNode) (head *ListNode) {
	var tail *ListNode
	carry := 0
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		sum := n1 + n2 + carry
		sum, carry = sum%10, sum/10
		if head == nil {
			head = &ListNode{Val: sum}
			tail = head
		} else {
			tail.Next = &ListNode{Val: sum}
			tail = tail.Next
		}
	}
	if carry > 0 {
		tail.Next = &ListNode{Val: carry}
	}
	return
}

func main() {
	//a := []int{2,4, 3}
	//b := []int{5,6,4}

	var a, b *ListNode

	ret := addTwoNumbers(a, b)
	fmt.Println(ret)

}


// 空间借用, 节省空间容量
func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	h := l1
	z := 0
	x := h
	for !(l1 == nil && l2 == nil) {
		a := 0
		b := 0
		if l1 != nil {
			a = l1.Val
		}
		if l2 != nil {
			b = l2.Val
		}
		x.Val = (a + b + z) % 10
		z = (a + b + z) / 10
		if (l2 == nil || l1 == nil) && z == 0 {
			break
		}
		if x.Next == nil && l2 != nil {
			l1 = nil
			x.Next = l2.Next
		}
		if x.Next == nil && z != 0 {
			x.Next = &ListNode{Val: z}
			break
		}
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
		x = x.Next
	}
	return h
}