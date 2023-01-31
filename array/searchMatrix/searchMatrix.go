package searchmatrix

/*
	https://leetcode.cn/problems/search-a-2d-matrix/
*/

func searchMatrix(matrix [][]int, target int) bool {
	var binarySearch func(nums []int,i,j int,target int) int
	binarySearch = func(nums []int, i, j int,target int) int {
		for i < j {
			mid := i + (j-i)/2
			if nums[mid] < target {
				i = mid+1
			} else {
				j = mid
			}
		}
		return i
	}
	var arr []int
	for i := range matrix {
		arr = append(arr, matrix[i]...)
	}
	resIndex := binarySearch(arr,0,len(arr),target)
	if resIndex < len(arr) && arr[resIndex] == target {
		return true
	}
	return false
}