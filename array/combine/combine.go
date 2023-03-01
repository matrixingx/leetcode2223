package combine

func combine(n int, k int) [][]int {
	var res [][]int
	var ans []int
	var dir []int
	for i := 1 ; i <= n ; i ++ {
		dir = append(dir, i)
	}
	var dfs func(index int)
	dfs = func(index int) {
		if len(ans) == k {
			res = append(res, append([]int{}, ans...))
			return
		}
		if index >= n || len(ans) > k{
			return
		}
		dfs(index+1)
		ans = append(ans, dir[index])
		dfs(index+1)
		ans = ans[:len(ans)-1]
	}
	dfs(0)
	return res
}