package partition

/*
	https://leetcode.cn/problems/partition-list/
*/

type ListNode struct {
	Val int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	var small,large = &ListNode{},&ListNode{}
	smallHead := small
	largeHead := large
	for head != nil {
		if head.Val < x {
			small.Next = head
			small = small.Next
		} else {
			large.Next = head
			large = large.Next
		}
		head = head.Next
	}
	small.Next = largeHead.Next
	large.Next = nil
	return smallHead.Next
}


// 错误版本，用快排思路不能保证比x大的数字与比x小的数字的相对顺序
func partition2(head *ListNode, x int) *ListNode {
	var arr []int
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	var i ,j = 0,len(arr)-1
	for i <= j {
		for arr[i] < x {
			i++
		}
		for arr[j] >= x {
			j--
		}
		if i <= j {
			arr[i],arr[j] = arr[j],arr[i]
			i++
			j--
		}
	}
	return mkListNode(arr)
}

func mkListNode(arr []int) *ListNode {
	var prev *ListNode
	for j := len(arr)-1 ; j >= 0 ; j-- {
		var temp = &ListNode{
			Val: arr[j],
			Next: prev,
		}
		prev = temp
	}
	return prev
}