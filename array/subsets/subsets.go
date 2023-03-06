package subsets

func subsets(nums []int) [][]int {
	var n = len(nums)
	var res [][]int
	var dfs func(index int,ans []int)
	dfs = func(index int, ans []int) {
		if index == n {
			res = append(res, append([]int{}, ans...))
			return
		}
		dfs(index+1,ans)
		ans = append(ans, nums[index])
		dfs(index+1,ans)
		ans = ans[:len(ans)-1]
	}
	dfs(0,[]int{})
	return res
}