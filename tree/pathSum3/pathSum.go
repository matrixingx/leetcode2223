package pathsum3

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	var res = 0
	res += rootSum(root,targetSum)
	res += pathSum(root.Left,targetSum)
	res += pathSum(root.Right,targetSum)
	return res
}

func rootSum (node *TreeNode, target int) int {
	var res = 0
	if node == nil {
		return res
	}
	if node.Val == target {
		res++
	}
	res += rootSum(node.Left,target-node.Val)
	res += rootSum(node.Right,target-node.Val)
	return res
}	