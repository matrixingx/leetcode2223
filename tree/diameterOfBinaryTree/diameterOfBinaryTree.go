package diameterofbinarytree

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	var dfs func(root *TreeNode) int
	var res = 1
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := dfs(root.Left)
		right := dfs(root.Right)
		res = max(res,left+right+1)
		return max(left,right)+1
	}
	dfs(root)
	return res -1
}

func max(i,j int) int {
	if i > j {
		return i
	}
	return j
}