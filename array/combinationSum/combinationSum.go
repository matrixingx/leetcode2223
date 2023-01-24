package combinationSum

/*
	https://leetcode.cn/problems/combination-sum/
*/

func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	var n = len(candidates)
	var backTrack func(sum,startIndex int,ans []int)
	backTrack = func(sum,startIndex int, ans []int) {
		if sum < 0 {
			return
		}
		if sum == 0 {
			res = append(res, append([]int{}, ans...))
			return
		}
		for i := startIndex ; i < n ; i++ {
			v := sum - candidates[i]
			if v >= 0 {
				ans = append(ans, candidates[i])
				backTrack(v,i,ans)
				ans = ans[:len(ans)-1]
			}
		}
	}
	backTrack(target,0,[]int{})
	return res
}