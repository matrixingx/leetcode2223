package deleteduplicates

/*
	https://leetcode.cn/problems/remove-duplicates-from-sorted-list-ii/
*/

type ListNode struct {
	Val int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var p = &ListNode{}
	p.Next = head
	var res = p
	for p.Next != nil && p.Next.Next != nil{
		if p.Next.Val == p.Next.Next.Val {
			v := p.Next.Val
			for p.Next != nil && v == p.Next.Val {
				p.Next = p.Next.Next
			}
		} else {
			p = p.Next
		}
	}
	return res.Next
}