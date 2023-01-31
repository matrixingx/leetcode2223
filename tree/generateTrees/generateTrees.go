package generatetrees

/*
	https://leetcode.cn/problems/unique-binary-search-trees-ii/
*/

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	var gen func(start,end int) []*TreeNode
	gen = func(start, end int) []*TreeNode {
		if start > end {
			return []*TreeNode{nil}
		}
		res := []*TreeNode{}
		for i := start ; i <= end ; i ++ {
			// i 代表以i为根
			leftTrees := gen(start,i-1)
			rightTrees := gen(i+1,end)
			for _,left := range leftTrees {
				for _,right := range rightTrees {
					root := TreeNode{
						Val:i,
						Left: left,
						Right: right,
					}
					res = append(res, &root)
				}
			}
		}
		return res
	}
	if n == 0 {
		return nil
	}
	return gen(1,n)
}