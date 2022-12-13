package flatten

/*
	https://leetcode.cn/problems/flatten-binary-tree-to-linked-list/
*/

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode)  {
	if root == nil {
		return
	}
	cur := root
	for cur != nil {
		if cur.Left != nil {
			p := cur.Left
			for p.Right != nil {
				p = p.Right
			}
			p.Right = cur.Right
			cur.Left,cur.Right = nil,cur.Left
		}
		cur = cur.Right
	}
}