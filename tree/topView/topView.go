package topview

import "sort"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ValWithDir struct {
	Val int
	Dir int
}

func topView(root *TreeNode) []int {
	var res []int
	var dirMap = make(map[int]ValWithDir)
	var dfs func(node *TreeNode, leftCount, rightCount int)
	dfs = func(node *TreeNode, leftCount, rightCount int) {
		if node == nil {
			return
		}
		if _, ok := dirMap[leftCount-rightCount]; !ok {
			dirMap[leftCount-rightCount] = ValWithDir{
				Val: node.Val,
				Dir: leftCount - rightCount,
			}
		}
		dfs(node.Left, leftCount+1, rightCount)
		dfs(node.Right, leftCount, rightCount+1)
	}
	dfs(root, 0, 0)
	var arr []ValWithDir
	for _, v := range dirMap {
		arr = append(arr, v)
	}
	sort.SliceStable(arr,func(i, j int) bool {
		return arr[i].Dir>arr[j].Dir
	})
	for i := range arr {
		res = append(res, arr[i].Val)
	}
	return res
}