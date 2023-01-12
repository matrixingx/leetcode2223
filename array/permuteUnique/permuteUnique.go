package permuteunique

import "sort"

func permuteUnique(nums []int) [][]int {
	var res [][]int
	var n = len(nums)
	var dfs = func(idx int,arr []int) {}
	sort.Ints(nums)
	var visited = make([]bool,n)
	dfs = func(idx int,arr []int) {
		if idx == n {
			res = append(res, append([]int{}, arr...))
			return
		}
		for i,v := range nums {
			if visited[i] || (i > 0 && !visited[i-1] && nums[i-1] == v) {
				continue
			}
			arr = append(arr, v)
			visited[i] = true
			dfs(idx+1,arr)
			visited[i] = false
			arr = arr[:len(arr)-1]
		}
	}
	dfs(0,[]int{})
	return res
}
