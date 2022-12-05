package reverseBetween

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
	Val: 5,
	Next: nil,
}

var testNode3 = &ListNode{
	Val: 1,
	Next: &ListNode{
		Val: 2,
		Next:nil,
	},
}

var testNode4 = &ListNode{
	Val: 1,
	Next: &ListNode{
		Val: 2,
		Next:&ListNode{
			Val: 3,
			Next:nil,
		},
	},
}

func TestReverseBetween(t *testing.T) {
	res := reverseBetween(testNode,2,4)
	fmt.Println(res)
	res = reverseBetween(testNode2,1,1)
	fmt.Println(res)
	res = reverseBetween(testNode3,1,2)
	fmt.Println(res)
	res = reverseBetween(testNode4,1,2)
	fmt.Println(res)
}

func TestReverseList(t *testing.T) {
	reverseHead,reverseTail := reverseList(testNode,3)
	fmt.Println(reverseHead)
	fmt.Println(reverseTail)
}