package reversekgroup

type ListNode struct {
	Val int
	Next *ListNode
}

func reverseList(node *ListNode) *ListNode {
	var prev *ListNode
	for node != nil {
		next := node.Next
		node.Next = prev
		prev = node
		node = next
	}
	return prev
}

func reverseList2(node *ListNode) *ListNode {
	if node == nil || node.Next == nil {
		return node
	}
	newHead := reverseList2(node.Next)
	node.Next.Next = node
	node.Next = nil
	return newHead
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	var dummyNode = &ListNode{Next: head}
	var prev,cur = dummyNode,dummyNode
	for cur != nil {
		for i := 0 ; i < k && cur != nil ; i++ {
			cur = cur.Next
		}
		if cur == nil {
			continue
		}
		next := cur.Next
		cur.Next = nil
		res := reverseList(prev.Next)
		prev.Next = res
		var temp = res
		for temp.Next != nil {
			temp = temp.Next
		}
		temp.Next = next
		prev,cur = temp,temp
	}
	return dummyNode.Next
}