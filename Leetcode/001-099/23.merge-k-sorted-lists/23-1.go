package _3_merge_k_sorted_lists

type ListNode struct {
	Val  int
	Next *ListNode
}

// leetcode book
// 与 21 类似; 此处使用分治思想
func mergeKLists(lists []*ListNode) *ListNode {
	lenth := len(lists)
	if lenth < 1 {
		return nil
	}
	if lenth == 1 {
		return lists[0]
	}

	mid := lenth / 2
	left := mergeKLists(lists[0:mid])
	right := mergeKLists(lists[mid:])

	return mergeTwoLists(left, right)
}

func mergeTwoLists(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2

	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	}
	l2.Next = mergeTwoLists(l1, l2.Next)
	return l2
}
