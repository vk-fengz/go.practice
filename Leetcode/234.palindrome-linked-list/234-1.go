package _34_palindrome_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

// 快慢指针查找中点, 反转链表
func isPalindrome(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	// 链表chain元素个数 ==> 奇数
	// 最后一个元素
	if fast != nil {
		slow = slow.Next
	}
	left := head
	right := reverse(slow)
	for right != nil {
		if left.Val != right.Val {
			return false
		}
		left = left.Next
		right = right.Next
	}

	return true
}

func reverse(head *ListNode) *ListNode {
	var pre *ListNode
	cur, next := head, head

	for cur != nil {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

// 其他解法: https://books.halfrost.com/leetcode/ChapterFour/0200~0299/0234.Palindrome-Linked-List/
// 空间: 复制一份 进行比较
// 快慢指针: 翻转后半部分, 进行对比;
