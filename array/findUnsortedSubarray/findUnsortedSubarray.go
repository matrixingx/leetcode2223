package findunsortedsubarray


func findUnsortedSubarray(nums []int) int {
	var n = len(nums)
	var leftMax, rightMin = -1 << 31, 1 << 31
	var leftIndex, rightIndex = -1, -1
	for i := range nums {
		if nums[i] < leftMax {
			rightIndex = i
		} else {
			leftMax = nums[i]
		}
		if nums[n-1-i] > rightMin {
			leftIndex = n - 1 - i
		} else {
			rightMin = nums[n-1-i]
		}
	}
	if rightIndex == -1{
		return 0
	}
	return rightIndex - leftIndex + 1
}