package binarysearch2

func search(nums []int, target int) bool {
	var binarysearch func(i,j int) int 
	var n = len(nums)
	binarysearch = func(i, j int) int {
		for i <= j {
			mid := i + (j-i)/2
			if nums[mid] == target {
				return mid
			}
			if nums[i] == nums[mid] && nums[mid] == nums[j] {
				i++
				j--
			} else if nums[i] <= nums[mid] { // 左边有序
				if target <= nums[mid] && target >= nums[i] {
					j = mid-1
				} else {
					i = mid+1
				}
			} else { // 右边有序
				if target <= nums[j] && target >= nums[mid] {
					i = mid+1
				} else {
					j = mid-1
				}
			}
		}
		return i
	}
	i := binarysearch(0,n-1)
	if i < n && nums[i] == target {
		return true
	}
	return false
}