package codec

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	var dfs func(node *TreeNode)
	var res = strings.Builder{}
	dfs = func(node *TreeNode) {
		if node == nil {
			res.WriteString("null,")
			return
		}
		v := strconv.Itoa(node.Val)
		res.WriteString(v)
		res.WriteString(",")
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return res.String()
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	s := strings.Split(data,",")
	s = s[:len(s)-1]
	var build func() *TreeNode
	build = func() *TreeNode {
		if s[0] == "null" {
			s = s[1:]
			return nil
		}
		v,_ := strconv.Atoi(s[0])
		s = s[1:]
		left := build()
		right := build()
		return &TreeNode{
			Val: v,
			Left: left,
			Right: right,
		}
	}
	return build()
}
