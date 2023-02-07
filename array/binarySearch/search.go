package binarysearch

func search(nums []int, target int) int {
	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		}
		return -1
	}
	var n = len(nums)
	var i,j = 0,n-1
	for i <= j {
		mid := (i+j)/2
		if nums[mid] == target {
			return mid
		}
		if nums[0] <= nums[mid] { // 如果左边有序
			if nums[0] <= target && target < nums[mid] { // 如果target在有序的左边
				j = mid-1
			} else {
				i = mid+1
			}
		} else { // 如果右边有序
			if nums[mid] <= target && target <= nums[n-1] { // 如果target在有序右边
				i = mid+1
			} else {
				j = mid-1
			}
		}
	}
	return -1
}