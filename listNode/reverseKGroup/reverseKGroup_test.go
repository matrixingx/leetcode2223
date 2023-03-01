package reversekgroup

import "testing"

func TestReverseKGroup(t *testing.T) {
	var node = &ListNode{
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
	reverseKGroup(node,2)
}