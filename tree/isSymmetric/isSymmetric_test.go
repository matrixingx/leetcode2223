package issymmetric

import (
	"fmt"
	"testing"
)

func TestIsSymmetric(t *testing.T) {
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
	}
	fmt.Println(isSymmetric(tree))
}