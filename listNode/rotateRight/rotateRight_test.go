package rotateright

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
					Next: nil,
				},
			},
		},
	},
}

var testNode2 = &ListNode{
	Val: 0,
	Next: &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: nil,
		},
	},
}

var testNode3 =  &ListNode{
	Val: 1,
	Next: &ListNode{
		Val: 2,
		Next: nil,
	},	
}

func TestRotateright(t *testing.T) {
	res := rotateRight(testNode,2)
	fmt.Println(res)
	res = rotateRight(testNode2,4)
	fmt.Println(res)
	res = rotateRight(testNode3,1)
	fmt.Println(res)
}