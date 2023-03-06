package maxslidingwindow

func maxSlidingWindow(nums []int, k int) []int {
	var n = len(nums)
	var queue []int
	var push func(i int)
	var res []int
	push = func(i int) {
		v := nums[i]
		for len(queue) > 0 && v >= nums[queue[len(queue)-1]] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, i)
	}
	for i := 0 ; i < k ; i ++ {
		push(i)
	}
	res = append(res, nums[queue[0]])
	for i := k ; i < n ; i ++ {
		push(i)
		for len(queue) > 0 && i-k >= queue[0] {
			queue = queue[1:]
		}
		res = append(res, nums[queue[0]])
	}
	return res
}