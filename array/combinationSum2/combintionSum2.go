package combinationsum2

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	var res [][]int
	var n = len(candidates)
	var visited = make([]bool,n)
	var backTrack func(sum, startIndex int, ans []int)
	sort.SliceStable(candidates, func(i, j int) bool {
		return candidates[i] < candidates[j]
	})
	backTrack = func(sum, startIndex int, ans []int) {
		if sum < 0 {
			return
		}
		if sum == 0 {
			res = append(res, append([]int{}, ans...))
			return
		}
		for i := startIndex; i < n; i++ {
			if i>0 && !visited[i-1] && candidates[i-1] == candidates[i] {
				continue
			}
			v := sum - candidates[i]
			if v >= 0 {
				ans = append(ans, candidates[i])
				visited[i] = true
				backTrack(v, i+1, ans)
				visited[i] = false
				ans = ans[:len(ans)-1]
			}
		}
	}
	backTrack(target, 0, []int{})
	return res
}
