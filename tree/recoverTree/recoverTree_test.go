package recovertree

import (
	"testing"
)

var root = &TreeNode{
	Val:1,
	Left: &TreeNode{
		Val:3,
		Left: nil,
		Right: &TreeNode{
			Val: 2,
		},
	},
}

func TestRecoverTree(t *testing.T) {
	recoverTree(root)
}