package sumnumbers

/*
	https://leetcode.cn/problems/sum-root-to-leaf-numbers/
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	var dfs func(node *TreeNode,sum int)
	var res int
	dfs = func(node *TreeNode,sum int) {
		if node == nil {
			return
		}
		if node.Left == nil && node.Right == nil{
			res += sum*10 + node.Val
			return
		}
		dfs(node.Left,sum*10 + node.Val)
		dfs(node.Right,sum*10 + node.Val)
	}
	dfs(root,0)
	return res
}