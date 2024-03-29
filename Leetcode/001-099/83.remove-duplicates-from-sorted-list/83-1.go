package _3_remove_duplicates_from_sorted_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	slow, fast := head, head
	for fast != nil {
		if slow.Val != fast.Val {
			slow.Next = fast
			slow = slow.Next
		} else {
			fast = fast.Next
		}
	}
	slow.Next = nil
	return head
}

// gobook-leetcode 采用了好多新的思想 来解决; 迭代的方法
