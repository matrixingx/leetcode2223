package jump

func jump(nums []int) int {
	var n = len(nums)
	var start = 0
	var jumpCount = 0
	if n == 1 {
		return jumpCount
	}
	for start < n {
		if start + nums[start] >= n-1 {
			jumpCount++
			break
		}
		nextMax := 0
		nextIndex := 0
		for i := 0 ; i <= nums[start] ; i++ {
			if start + i + nums[start + i] > nextMax {
				nextMax = start + i + nums[start + i]
				nextIndex = start + i
			}
		}
		start = nextIndex
		jumpCount++
	}
	return jumpCount
}