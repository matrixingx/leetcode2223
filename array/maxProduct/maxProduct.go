package maxproduct

func maxProduct(nums []int) int {
	var n = len(nums)
	if n == 0 {
		return 0
	}
	var minDp,maxDP = make([]int,n),make([]int,n)
	var ans = 0
	minDp[0]=nums[0]
	maxDP[0]=nums[0]
	ans = nums[0]
	for i := 1 ; i < n ; i ++ {
		maxDP[i] = max(maxDP[i-1]*nums[i],max(nums[i],minDp[i-1]*nums[i]))
		minDp[i] = min(minDp[i-1]*nums[i],min(nums[i],maxDP[i-1]*nums[i]))
		ans = max(maxDP[i],ans)
	}
	return ans
}

func max(i,j int) int {
	if i > j {
		return i
	}
	return j
}

func min(i,j int) int {
	if i < j {
		return i
	}
	return j
}