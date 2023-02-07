package pathsum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	var dfs func(node *TreeNode,arr []int,target int)
	var res [][]int
	dfs = func(node *TreeNode, arr []int,target int) {
		if node == nil {
			return
		}
		target -= node.Val
		arr = append(arr, node.Val)
		dfs(node.Left,arr,target)
		dfs(node.Right,arr,target)
		if target == 0 && node.Left == nil && node.Right == nil {
			res = append(res, append([]int{},arr...))
		}
		target+=node.Val
		arr = arr[:len(arr)-1]
	}
	dfs(root,[]int{},targetSum)
	return res
}