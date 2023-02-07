package sumnumbers

import (
	"fmt"
	"testing"
)

var tree = &TreeNode{
	Val: 4,
	Left: &TreeNode{
		Val: 9,
		Left: &TreeNode{
			Val: 5,
		},
		Right: &TreeNode{
			Val: 1,
		},
	},
	Right: &TreeNode{
		Val: 0,
	},
}

func TestSumNumbers(t *testing.T) {
	fmt.Println(sumNumbers(tree))
}
