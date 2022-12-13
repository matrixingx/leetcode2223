package detectcycle

import (
	"fmt"
	"testing"
)

var testNode = &ListNode{
	Val:  3,
	Next: testNode2,
}

var testNode2 = &ListNode{
	Val: 2,
	Next: &ListNode{
		Val:  0,
		Next: testNode3,
	},
}

var testNode3 = &ListNode{
	Val:  -4,
	Next: nil,
}

func TestDetectCycle(t *testing.T) {
	//testNode3.Next = testNode2
	res := detectCycle(testNode)
	if res != nil {
		fmt.Println(res.Val)
	}
}