package connect2

type LevelType int64

type Node struct {
	Val int
	Left *Node
	Right *Node
	Next *Node
}

func connect(root *Node) *Node {
	var res = root
	var p = root
	if p == nil {
		return p
	}
	for p != nil  {
		for levelP := p ; levelP != nil ; levelP = levelP.Next {
			if levelP.Left != nil {
				levelP.Left.Next = getLeftNext(levelP)
			}
			if levelP.Right != nil {
				levelP.Right.Next = getRightNext(levelP)
			}
		}
		for p != nil {
			if p.Left != nil {
				p = p.Left
				break
			} else if p.Right != nil {
				p = p.Right
				break
			} else {
				p = p.Next
			}
		}
	}
	return res
}

func getLeftNext(node *Node) *Node {
	if node.Left != nil {
		if node.Right != nil {
			return node.Right
		}
	}
	for p := node.Next;p != nil ; p = p.Next {
		if p.Left != nil {
			return p.Left
		}
		if p.Right != nil {
			return p.Right
		}
	}
	return nil
}

func getRightNext(node *Node) *Node {
	for p := node.Next;p != nil ; p = p.Next {
		if p.Left != nil {
			return p.Left
		}
		if p.Right != nil {
			return p.Right
		}
	}
	return nil
}