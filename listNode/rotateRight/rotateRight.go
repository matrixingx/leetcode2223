package rotateright

/*
	https://leetcode.cn/problems/rotate-list/
*/

type ListNode struct {
	Val int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k == 0{
		return head
	}
	var arr []int
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	k = k % len(arr)
	arr = append(arr[len(arr)-k : len(arr)], arr[:len(arr)-k]...)
	return mkListNode(arr)
}

func mkListNode(arr []int) *ListNode {
	var prev *ListNode
	for j := len(arr)-1 ; j >= 0 ; j -- {
		temp := &ListNode{
			Val: arr[j],
			Next: prev,
		}
		prev = temp
	}
	return prev
}