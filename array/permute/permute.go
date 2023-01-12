package permute

func permute(nums []int) [][]int {
	var res [][]int
	var n = len(nums)
	var dfs = func(arr []int) {}
	dfs = func(arr []int) {
		if len(arr) == n {
			res = append(res, append([]int{}, arr...))
			return
		}
		for i := range nums {
			v := nums[i]
			nums = append(nums[:i], nums[i+1:]...)
			arr = append(arr, v)
			dfs(arr)
			arr = arr[:len(arr)-1]
			nums = append(nums[:i], append([]int{v}, nums[i:]...)...)
		}
	}
	dfs([]int{})
	return res
}
