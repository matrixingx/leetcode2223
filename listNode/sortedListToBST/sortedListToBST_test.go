package sortedlisttobst

import (
	"fmt"
	"testing"
)

var testList = &ListNode{
	Val: -10,
	Next:&ListNode{
		Val: -3,
		Next:&ListNode{
			Val: 0,
			Next:&ListNode{
				Val: 5,
				Next:&ListNode{
					Val: 9,
					Next:nil,
				},
			},
		},
	},
}

func TestSortedListToBST(t *testing.T) {
	testArr := []int{-10,-3,0,5,9}
	idx := len(testArr)/2
	lft := testArr[:idx]
	rgt := testArr[idx+1:]
	fmt.Println(lft)
	fmt.Println(rgt)
	res := sortedListToBST(testList)
	fmt.Println(res.Val)
}