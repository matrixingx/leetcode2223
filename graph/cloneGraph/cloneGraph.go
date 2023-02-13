package clonegraph

type Node struct {
	Val int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
    var visited = make(map[*Node]*Node)
	var copy func(n *Node) *Node 
	copy = func(n *Node) *Node {
		if n == nil {
			return n
		}
		if vis,ok := visited[n];ok {
			return vis
		}
		newNode := &Node{
			Val: n.Val,
			Neighbors: []*Node{},
		}
		visited[n] = newNode
		for _,v := range n.Neighbors {
			newNode.Neighbors = append(newNode.Neighbors, copy(v))
		}
		return newNode
	}
	return copy(node)
}