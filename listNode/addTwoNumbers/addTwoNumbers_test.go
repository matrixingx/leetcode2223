package addtwonumbers

import (
	"fmt"
	"testing"
)

var testNode = &ListNode{
	Val: 9,
	Next: &ListNode{
		Val: 9,
		Next: &ListNode{
			Val: 9,
			Next: &ListNode{
				Val: 9,
				Next: &ListNode{
					Val: 9,
					Next: &ListNode{
						Val: 9,
						Next: &ListNode{
							Val:  9,
							Next: nil,
						},
					},
				},
			},
		},
	},
}

var testNode2 = &ListNode{
	Val: 9,
	Next: &ListNode{
		Val: 9,
		Next: &ListNode{
			Val: 9,
			Next: &ListNode{
				Val:  9,
				Next: nil,
			},
		},
	},
}

var testNode3 = &ListNode{
	Val: 2,
	Next: &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 3,
		},
	},
}

var testNode4 = &ListNode{
	Val: 5,
	Next: &ListNode{
		Val: 6,
		Next: &ListNode{
			Val: 4,
		},
	},
}

func TestAddTwoNumbers(t *testing.T) {
	res := addTwoNumbers(testNode,testNode2)
	for res != nil {
		fmt.Printf("%v",res.Val)
		res = res.Next
	}
	fmt.Println()
	res = addTwoNumbers(testNode3,testNode4)
	for res != nil {
		fmt.Printf("%v",res.Val)
		res = res.Next
	}
	fmt.Println()
}