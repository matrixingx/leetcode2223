package test6

type ListNode struct {
	Val int
	Next *ListNode
}

var head = &ListNode{
	Val:4,
	Next: &ListNode{
		Val:3,
		Next: &ListNode{
			Val:2,
			Next: &ListNode{
				Val:1,
				Next: nil,
			},
		},
	},
}

func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var dummy = &ListNode{Next: head}
	var sorted,cur = head,head.Next
	for cur != nil {
        if cur.Val >= sorted.Val {
            sorted = sorted.Next
        } else {
            temp := dummy
            for temp != nil && temp.Next != nil && temp.Next.Val < cur.Val {
                temp = temp.Next
            }
            sorted.Next = cur.Next
            cur.Next = temp.Next
            temp.Next = cur
            cur = sorted
        }
        cur = cur.Next
	}
	return dummy.Next
}