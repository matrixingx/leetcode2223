package canjump

func canJump(nums []int) bool {
	var n = len(nums)
	var start = 0
	for start < n {
		j := nums[start]
		if start + j >= n-1 {
			return true
		}
		if j == 0 {
			return false
		}
		// 贪心寻找下次能跳最远的点
		nextMax := 0
		nextIndex := 0
		for k := 0 ; k <= j ; k++ {
			if k + nums[start+k] >= nextMax {
				nextMax = k + nums[start+k]
				nextIndex = start+k
			}
		}
		if nextMax >= n-1 {
			break
		}
		start = nextIndex
	}
	return true
}