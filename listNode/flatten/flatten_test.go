package flatten

import (
	"fmt"
	"testing"
)

var testTree = &TreeNode{
	Val: 1,
	Left: nil,
	Right: &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 3,
		},
		Right: nil,
	},
}

func TestFlatten(t *testing.T) {
	flatten(testTree)
	fmt.Println()
}