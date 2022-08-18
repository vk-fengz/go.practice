package _6_partition_lis

type ListNode struct {
	Val  int
	Next *ListNode
}

// > https://books.halfrost.com/leetcode/ChapterFour/0001~0099/0086.Partition-List/

func partition(head *ListNode, x int) *ListNode {
	beforeHead := &ListNode{Val: 0, Next: nil}
	afterHead := &ListNode{Val: 0, Next: nil}
	before, after := beforeHead, afterHead

	for head != nil {
		if head.Val < x {
			before.Next = head
			before = before.Next
		} else if head.Val >= x {

			after.Next = head
			after = after.Next
		}
		head = head.Next
	}
	after.Next = nil
	before.Next = afterHead.Next
	return beforeHead.Next
}
