package _9_remove_nth_node_from_end_of_list

type ListNode struct {
	Val  int
	Next *ListNode
}

// 虚拟头结点的使用; 避免删除头结点, 防止出现空指针
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{Next: head}

	slowpre, slow, fast := dummyHead, head, head
	for fast != nil {
		if n < 1 {
			slowpre = slow
			slow = slow.Next
		}
		n--
		fast = fast.Next
	}
	slowpre.Next = slow.Next
	return dummyHead.Next
}
