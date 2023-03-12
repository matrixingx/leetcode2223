package subarraysum

func subarraySum(nums []int, k int) int {
	var subMp = make(map[int]int)
	var res int
	var n = len(nums)
	var sum = 0
	subMp[0] = 1
	for i := 0 ; i < n ; i++ {
		sum += nums[i]
		if v,ok := subMp[sum-k];ok {
			res += v
		}
		subMp[sum]++
	}
	return res
}