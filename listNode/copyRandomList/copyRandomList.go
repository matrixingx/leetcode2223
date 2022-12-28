package copyrandomlist

type Node struct {
	Val int
	Next *Node
	Random *Node
}
/*
	https://leetcode.cn/problems/copy-list-with-random-pointer/
*/
// 映射关系：旧链表节点->新链表节点
var cacheMap = make(map[*Node]*Node)

func copyRandomList(head *Node) *Node {
	return deepCopy(head)
}

func deepCopy(node *Node) *Node {
	if node == nil {
		return nil
	}
	if n,ok := cacheMap[node];ok {
		return n
	}
	newNode := &Node{
		Val:node.Val,
	}
	cacheMap[node] = newNode
	newNode.Next = deepCopy(node.Next)
	newNode.Random = deepCopy(node.Random)
	return newNode
}