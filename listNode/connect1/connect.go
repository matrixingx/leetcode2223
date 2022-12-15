package connect1

/*
	https://leetcode.cn/problems/populating-next-right-pointers-in-each-node/
*/

type Node struct {
	Val int
	Left *Node
	Right *Node
	Next *Node
}

func connect(root *Node) *Node {
	var res = root
	if root == nil {
		return root
	}
	for p := root ; p.Left != nil ; p = p.Left {
		for levelP := p;levelP != nil;levelP = levelP.Next {
			levelP.Left.Next = levelP.Right
			if levelP.Next != nil {
				levelP.Right.Next = levelP.Next.Left
			}
		}
	}
	return res
}