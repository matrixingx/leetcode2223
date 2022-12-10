package removenthfromend

import (
	"fmt"
	"testing"
)

var testNode = &ListNode{
	Val: 1,
	Next: &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 5,
				},
			},
		},
	},
}

var testNode2 = &ListNode{
	Val: 1,
}

var testNode3 = &ListNode{
	Val: 1,
	Next: &ListNode{
		Val: 2,
	},
}


func TestRemoveNthFromEnd(t *testing.T) {
	res := removeNthFromEnd(testNode,2)
	for res != nil {
		fmt.Printf("%v",res.Val)
		res = res.Next
	}
	fmt.Println()
	res = removeNthFromEnd(testNode2,1)
	for res != nil {
		fmt.Printf("%v",res.Val)
		res = res.Next
	}
	fmt.Println()
	res = removeNthFromEnd(testNode3,1)
	for res != nil {
		fmt.Printf("%v",res.Val)
		res = res.Next
	}
	fmt.Println()
}