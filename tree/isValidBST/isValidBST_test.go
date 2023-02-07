package isvalidbst

import (
	"fmt"
	"testing"
)

var tree = &TreeNode{
	Val:4,
	Left: &TreeNode{
		Val:3,
		Left: &TreeNode{
			Val:1,
			Left: nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:2,
			Left: nil,
			Right: nil,
		},
	},
	Right: &TreeNode{
		Val:7,
		Left: &TreeNode{
			Val:6,
			Left: nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:8,
			Left: nil,
			Right: nil,
		},
	},
}

var tree2 = &TreeNode{
	Val:4,
	Left: &TreeNode{
		Val:2,
		Left: &TreeNode{
			Val:1,
			Left: nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:3,
			Left: nil,
			Right: nil,
		},
	},
	Right: &TreeNode{
		Val:7,
		Left: &TreeNode{
			Val:6,
			Left: nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:8,
			Left: nil,
			Right: nil,
		},
	},
}

var tree3 = &TreeNode{
	Val:2,
	Left: &TreeNode{
		Val:2,
	},
	Right: &TreeNode{
		Val:2,
	},
}

func TestIsValidBST(t *testing.T) {
	fmt.Println(isValidBST(tree))
	fmt.Println(isValidBST(tree2))
	fmt.Println(isValidBST(tree3))
}