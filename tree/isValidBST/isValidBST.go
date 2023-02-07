package isvalidbst


type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	var dfs func(node *TreeNode)
	var arr []int
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		arr = append(arr, node.Val)
		dfs(node.Right)
	}
	dfs(root)
	var n = len(arr)
	for i := 0 ; i < n-1 ; i++ {
		if arr[i] >= arr[i+1] {
			return false
		}
	}
	return true
}