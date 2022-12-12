package sortedlisttobst

/*
	https://leetcode.cn/problems/convert-sorted-list-to-binary-search-tree/
	go分割数组 分割点左开右闭 也就是 ) [
*/

type ListNode struct {
	Val int
	Next *ListNode
}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	var arr []int
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	return mkBST(arr)
}

func mkBST(arr []int) *TreeNode {
	if len(arr) < 1 {
		return nil
	}
	if len(arr) == 1 {
		return &TreeNode{
			Val: arr[0],
		}
	}
	idx := len(arr)/2
	return &TreeNode{
		Val: arr[idx],
		Left: mkBST(arr[:idx]),
		Right: mkBST(arr[idx+1:]),
	}
}
