package partition

import (
	"fmt"
	"testing"
)

var testNode = &ListNode{
	Val: 1,
	Next: &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 5,
					Next: &ListNode{
						Val: 2,
						Next: nil,
					},
				},
			},
		},
	},
}

func TestPartition(t *testing.T) {
	res := partition(testNode,3)
	fmt.Println(res)
}