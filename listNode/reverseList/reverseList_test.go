package reverselist

import "testing"

var testList = &ListNode{
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

func TestReverseList(t *testing.T) {
	reverseList(testList)
}