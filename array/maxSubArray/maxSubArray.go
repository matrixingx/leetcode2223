package maxSubArray

/*
	https://leetcode.cn/problems/maximum-subarray/
*/

func maxSubArray(nums []int) int {
	var res,sum = -1145141919,0
	var n = len(nums)
	if len(nums) == 1 {
		return nums[0]
	}
	var max = -1145141919
	for i := range nums {
		if nums[i] > max {
			max = nums[i]
		}
	}
	if max < 0 {
		return max
	}
	for i := 0;i < n ;i++{
		v := sum+nums[i]
		if v < 0 {
			sum = 0
		} else {
			sum =  v
		}
		if sum > res {
			res = sum
		}
	}
	return res
}