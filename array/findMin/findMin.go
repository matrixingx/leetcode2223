package findmin

func findMin(nums []int) int {
	var binarySearch func(i,j int) int
	var start,end = 0,len(nums)-1
	md := start + (end-start)/2
	if nums[start] < nums[md] && nums[md] < nums[end] {// 特殊情况：有序
		return nums[start]
	}
	binarySearch = func(i, j int) int {
		for i < j {
			mid := i + (j-i)/2
			if nums[end] > nums[mid] { // mid左边有序，最小值再mid右边包括mid
				j = mid
			} else { // mid右边有序，最小值在mid左边
				i = mid+1
			}
		}
		return i
	}
	res := binarySearch(start,end)
	return nums[res]
}