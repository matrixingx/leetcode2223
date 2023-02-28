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
	for i := 0 ; i < left-1 ; i ++ {
		prev = prev.Next
	}
	var temp = prev.Next
	for i := 0 ; i < right - left ; i ++ {
		temp = temp.Next
	}
	var next = temp.Next
	temp.Next = nil
	res := reverseList(prev.Next)
	prev.Next = res
	for res.Next != nil {
		res = res.Next
	}
	res.Next = next
	return DummyNode.Next
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