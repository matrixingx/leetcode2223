package insertionsortlist

/*
	https://leetcode.cn/problems/insertion-sort-list/
*/

type ListNode struct {
	Val int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var dummy = &ListNode{Next: head}
	var sorted,cur = head,head.Next
	for cur != nil {
		if sorted.Val <= cur.Val {
			sorted = sorted.Next
		} else {
			prev := dummy
			for prev.Next.Val <= cur.Val {
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