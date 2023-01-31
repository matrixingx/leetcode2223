package recovertree

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func recoverTree(root *TreeNode)  {
	var inorder func(node *TreeNode)
	var arr []int
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		arr = append(arr, node.Val)
		inorder(node.Right)
		return
	}
	inorder(root)
	var n = len(arr)
	x,y := -1,-1
	for i := 0 ; i < n-1 ; i ++ {
		if arr[i]>arr[i+1] {
			y = i+1
			if x == -1 {
				x = i
			} else {
				break
			}
		}
	}
	var recover func(node *TreeNode,count int)
	recover = func(node *TreeNode, count int) {
		if node == nil || count == 0 {
			return
		}
		if count > 0 && (node.Val == arr[x] || node.Val == arr[y]){
			if node.Val == arr[x] {
				node.Val = arr[y]
			} else {
				node.Val = arr[x]
			}
			count--
		}
		recover(node.Left,count)
		recover(node.Right,count)
	}
	recover(root,2)
}