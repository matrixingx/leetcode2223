package removenthfromend

type ListNode struct {
	Val int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var p1,p2,res *ListNode
	res,p1,p2 = head,head,head
	for n > 0 && p2 != nil{
		p2 = p2.Next
		n--
	}
	if p2 == nil {
		return p1.Next
	}
	for p2.Next != nil {
		p2 = p2.Next
		p1 = p1.Next
	}
	if p1.Next != nil {
		p1.Next = p1.Next.Next
	}
	return res
}