package swappairs

/*
	https://leetcode.cn/problems/swap-nodes-in-pairs/
*/

type ListNode struct {
	Val int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	var p = &ListNode{}
	p.Next = head
	var res = p
	for p.Next != nil && p.Next.Next != nil {
		p1 := p.Next
		p2 := p.Next.Next
		p.Next = p2.Next
		p2.Next = p1
		p1.Next = p.Next
		p.Next = p2
		p = p.Next.Next
	}
	return res.Next
}