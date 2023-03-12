package codec

import (
	"fmt"
	"testing"
)

func TestCodec(t *testing.T) {
	testTree := &TreeNode{
		Val:1,
		Left: &TreeNode{
			Val:2,
			Left: nil,
			Right: &TreeNode{
				Val:3,
				Left: nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:4,
			Left: &TreeNode{
				Val:5,
				Left: nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:6,
				Left: nil,
				Right: &TreeNode{
					Val:7,
					Left: nil,
					Right: nil,
				},
			},
		},
	}
	c := Constructor()
	s := c.serialize(testTree)
	n := c.deserialize(s)
	fmt.Println(n.Val)
}