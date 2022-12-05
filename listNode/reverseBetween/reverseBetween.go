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
	var hd,reverse,tail,temp *ListNode
	temp = head
	for i := 1 ; i <= right + 1 ; i ++ {
		if temp == nil {
			break
		}
		if i == left - 1 {
			hd = temp
		}
		if i == left {
			reverse = temp
		}
		if i == right + 1 {
			tail = temp
		}
		temp = temp.Next
	}
	reverseHead,reverseTail := reverseList(reverse,right - left + 1)
	if hd != nil {
		hd.Next = reverseHead
	} else {
		head = reverseHead
	}
	if reverseTail != nil {
		reverseTail.Next = tail
	}
	return head
}

func reverseList(head *ListNode,n int) (prev,tail *ListNode) {
	tail = head
	for n > 0 {
		var temp,next *ListNode
		temp = head
		next = head.Next
		temp.Next = prev
		prev = temp
		head = next
		n--
	}
	return prev ,tail
}