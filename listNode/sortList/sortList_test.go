package sortlist

import (
	"fmt"
	"testing"
)

var testList = &ListNode{
	Val:-1,
	Next: &ListNode{
		Val:5,
		Next: &ListNode{
			Val:3,
			Next: &ListNode{
				Val:4,
				Next: &ListNode{
					Val:0,
					Next: &ListNode{
						Val:1,
						Next: nil,
					},
				},
			},
		},
	},
}

func TestSortList(t *testing.T) {
	res := sortList2(testList)
	fmt.Println(res.Val)
}