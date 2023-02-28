package lowestcommonancestor

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var dfs func(node *TreeNode) *TreeNode
	dfs =func(node *TreeNode) *TreeNode {
		if node == nil {
			return nil
		}
		if node == p {
			return p
		}
		if node == q {
			return q
		}
		left := dfs(node.Left)
		right := dfs(node.Right)
		if left != nil && right != nil {
			return node
		}
		if left != nil {
			return left
		}
		return right
	}
	res := dfs(root)
	return res
}