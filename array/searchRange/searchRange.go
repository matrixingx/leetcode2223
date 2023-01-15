package searchrange

func searchRange(nums []int, target int) []int {
	var n = len(nums)
	var binarySearch func(i, j int,t int) int
	// 左闭右开
	binarySearch = func(i, j int,t int) int {
		for i < j {
			mid := i + (j-i) / 2
			if nums[mid] < t {
				i = mid + 1
			} else {
				j = mid
			}
		}
		return i
	}
	start := binarySearch(0,n,target)
	if start == n || nums[start] != target {
		return []int{-1,-1}
	}
	end := binarySearch(start,n,target+1)-1
	return []int{start, end}
}
// 左右都闭
func binarySearch2(nums []int,target int) int {
	var i,j = 0,len(nums)-1
	for i <= j {
		mid := i + (j-i)/2
		if nums[mid] < target {
			i = mid+1
		} else { // it means nums[mid]>= target
			j = mid-1
		}
	}
	return i
}