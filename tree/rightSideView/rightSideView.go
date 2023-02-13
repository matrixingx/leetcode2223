package rightsideview

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	var res []int
	var q []*TreeNode
	q = append(q, root)
	var bfs func(queue []*TreeNode,node *TreeNode)
	bfs = func(queue []*TreeNode,node *TreeNode) {
		if node == nil {
			return
		}
		queue = append(queue, node)
		for len(queue) > 0 {
			n := len(queue)
			for i := 0 ; i < n ; i++ {
				tmpNode := queue[0]
				queue = queue[1:]
				if i == n-1 {
					res = append(res, tmpNode.Val)
				}
				if tmpNode.Left != nil {
					queue = append(queue, tmpNode.Left)
				}
				if tmpNode.Right != nil {
					queue = append(queue, tmpNode.Right)
				}
			}
		}
	}
	bfs(q,root)
	return res
}