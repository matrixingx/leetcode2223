package reverseBetween

/*
	https://leetcode.cn/problems/reverse-linked-list-ii/
*/

type ListNode struct {
	Val int
	Next *ListNode
}


func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == right {
		return head
	}
	DummyNode := &ListNode{Next: head}
	var prev = DummyNode
	for i := 0 ; i < left ; i
	return head
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev
}