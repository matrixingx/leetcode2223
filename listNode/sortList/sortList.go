package sortlist

type ListNode struct {
	Val int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	var dummy = &ListNode{Next: head}
	if head == nil || head.Next == nil {
		return head
	}
	var sorted,cur = head,head.Next
	for cur != nil {
		if sorted.Val < cur.Val {
			sorted = sorted.Next
		} else {
			prev := dummy
			for prev.Next.Val < cur.Val {
				prev = prev.Next
			}
			sorted.Next = cur.Next
			cur.Next = prev.Next
			prev.Next = cur
		}
		cur = sorted.Next
	}
	return dummy.Next
}

func sortList2(head *ListNode) *ListNode {
	var p = &ListNode{Next: head}
	var n = 0
	for p != nil {
		n++
		p = p.Next
	}
	var mergeList func(p *ListNode,length int) *ListNode
	var merge func(p,q *ListNode) *ListNode
	mergeList = func(p *ListNode, length int) *ListNode {
		if length < 2 {
			return p
		}
		mid := length / 2
		q := p
		for i := 0 ; i < mid ; i++ {
			if i == mid -1 {
				temp := q
				q = q.Next
				temp.Next = nil
				break
			}
			q = q.Next
		}
		left := mergeList(p,mid)
		right := mergeList(q,length-mid)
		res := merge(left,right)
		return res
	}
	merge = func(p, q *ListNode) *ListNode {
		var dummy = &ListNode{}
		var cur = dummy
		for p != nil && q != nil {
			if p.Val < q.Val {
				cur.Next = p
				p = p.Next
			} else {
				cur.Next = q
				q = q.Next
			}
			cur = cur.Next
		}
		for p != nil {
			cur.Next = p
			p = p.Next
			cur = cur.Next
		}
		for q != nil {
			cur.Next = q
			q = q.Next
			cur = cur.Next
		}
		return dummy.Next
	}
	return mergeList(head,n)
}