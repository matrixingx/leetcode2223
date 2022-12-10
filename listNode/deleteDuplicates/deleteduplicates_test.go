package deleteduplicates

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
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 5,
							Next: &ListNode{
								Val: 6,
							},
						},
					},
				},
			},
		},
	},
}

func TestDeleteDuplicates(t *testing.T) {
	res := deleteDuplicates(testNode)
	for res != nil {
		fmt.Printf("%v ",res.Val)
		res = res.Next
	}
	fmt.Println()
}