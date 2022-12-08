package addtwonumbers

/*
	https://leetcode.cn/problems/add-two-numbers/
*/

type ListNode struct {
	Val int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		resHead = &ListNode{
			Next: &ListNode{},
		}
		p = false
		hd *ListNode
	)
	hd = resHead
	for l1 != nil && l2 != nil {
		hd = hd.Next
		value := (l1.Val + l2.Val)
		if p {
			value ++
		}
		hd.Val = value % 10
		p = value / 10 == 1
		hd.Next = &ListNode{}
		l1 = l1.Next
		l2 = l2.Next
	}
	for l1 != nil {
		hd = hd.Next
		value := l1.Val
		if p {
			value ++
		}
		hd.Val = value % 10
		p = value / 10 == 1
		hd.Next = &ListNode{}
		l1 = l1.Next
	}
	for l2 != nil {
		hd = hd.Next
		value := l2.Val
		if p {
			value ++
		}
		hd.Val = value % 10
		p = value / 10 == 1
		hd.Next = &ListNode{}
		l2 = l2.Next
	}
	if p {
		hd.Next = &ListNode{
			Val: 1,
		}
	} else {
		hd.Next = nil
	}
	return resHead.Next
}