package rob

func rob(nums []int) int {
	var n = len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	if n == 2 {
		if nums[0] > nums[1] {
			return nums[0]
		} else {
			return nums[1]
		}
	}
	var dp = make([]int,n)
	dp[0] = nums[0]
	dp[1] = max(nums[0],nums[1])
	for i := 2 ; i < n ; i ++ {
		dp[i] = max(dp[i-2]+nums[i],dp[i-1])
	}
	return dp[n-1]
}

func max(i,j int) int {
	if i > j {
		return i
	}
	return j
}