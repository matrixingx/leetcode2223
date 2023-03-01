package issymmetric

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	var res = true
	var dfs func(node1,node2 *TreeNode)
	dfs = func(node1,node2 *TreeNode) {
		if !res {
			return
		}
		if node1 == nil && node2 == nil {
			return
		}
		if node1 == nil || node2 == nil {
			res = false
			return
		}
		if node1.Val != node2.Val {
			res = false
			return
		}
		dfs(node1.Left,node2.Right)
		dfs(node1.Right,node2.Left)
	}
	dfs(root,root)
	return res
}