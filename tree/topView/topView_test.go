package topview

import (
	"fmt"
	"testing"
)

func TestTopView(t *testing.T) {
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
				Left: nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val: 4,
				Left: nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 6,
				Right: &TreeNode{
					Val: 7,
					Left: nil,
					Right: &TreeNode{
						Val: 8,
						Left: nil,
						Right: nil,
					},
				},
			},
			
		},
	}
	fmt.Println(topView(tree))
}