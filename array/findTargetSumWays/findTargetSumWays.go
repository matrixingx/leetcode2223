package findtargetsumways

func findTargetSumWays(nums []int, target int) int {
	var backTrack func(index ,sum int)
	var n = len(nums)
	var res = 0
	backTrack = func(index, sum int) {
		if index == n {
			if sum == target{
				res++
			}
			return
		}
		backTrack(index+1,sum+nums[index])
		backTrack(index+1,sum-nums[index])
	}
	backTrack(0,0)
	return res
}