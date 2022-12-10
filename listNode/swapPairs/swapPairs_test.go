package swappairs

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
			},
		},
	},
}

var testNode2 = &ListNode{
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

func TestSwapPairs(t *testing.T) {
	hd := swapPairs(testNode)
	for hd != nil {
		fmt.Printf("%v ",hd.Val)
		hd = hd.Next
	}
	fmt.Println()
	hd = swapPairs(testNode2)
	for hd != nil {
		fmt.Printf("%v ",hd.Val)
		hd = hd.Next
	}
	fmt.Println()
}