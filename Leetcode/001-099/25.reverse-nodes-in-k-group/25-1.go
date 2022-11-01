package _5_reverse_nodes_in_k_group

type ListNode struct {
	Val  int
	Next *ListNode
}

//

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	a, b := head, head
	for i := 0; i < k; i++ {
		if b == nil {
			return head
		}
		b = b.Next
	}
	newHead := reverse(a, b)
	a.Next = reverseKGroup(b, k)

	return newHead
}

/** 反转区间 [a, b) 的元素，注意是左闭右开 */
func reverse(a, b *ListNode) *ListNode {
	var pre *ListNode
	cur, next := a, a

	for cur != b {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

// go-leetcodebook: https://books.halfrost.com/leetcode/ChapterFour/0001~0099/0025.Reverse-Nodes-in-k-Group/
