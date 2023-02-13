package rightsideview

import (
	"fmt"
	"testing"
)

var tree = &TreeNode{
	Val: 1,
	Left: &TreeNode{
		Val: 2,
		Right: &TreeNode{
			Val: 5,
		},
	},
	Right: &TreeNode{
		Val: 3,
		Right: &TreeNode{
			Val: 4,
		},
	},
}

var tree2 = &TreeNode{
	Val: 1,
	Left: &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 4,
		},
	},
	Right: &TreeNode{
		Val: 3,
	},
}

func TestRightSideView(t *testing.T) {
	fmt.Println(rightSideView(tree))
	fmt.Println(rightSideView(tree2))
}
