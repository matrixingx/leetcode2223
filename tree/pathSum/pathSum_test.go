package pathsum

import (
	"fmt"
	"testing"
)

var tree = &TreeNode{
	Val: 4,
	Left: &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   1,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   12,
			Left:  nil,
			Right: nil,
		},
	},
	Right: &TreeNode{
		Val: 7,
		Left: &TreeNode{
			Val:   6,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   8,
			Left:  nil,
			Right: nil,
		},
	},
}

var tree2 = &TreeNode{
	Val: -2,
	Right: &TreeNode{
		Val: -3,
	},
}

var tree3 = &TreeNode{
	Val: 1,
	Left: &TreeNode{
		Val: 2,
	},
}

func TestPathSum(t *testing.T) {
	fmt.Println(pathSum(tree, 17))
	fmt.Println(pathSum(tree, 19))
	fmt.Println(pathSum(tree2, -3))
	fmt.Println(pathSum(tree3, 0))
}
