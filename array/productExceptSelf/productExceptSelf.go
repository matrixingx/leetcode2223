package productexceptself

func productExceptSelf(nums []int) []int {
	var res []int
	var n = len(nums)
	var leftProduct = make([]int,n)
	var rightProduct = make([]int,n)
	for i := 0 ; i < n ; i++ {
		if i == 0 {
			leftProduct[i] = nums[i]
		} else {
			leftProduct[i] = nums[i]*leftProduct[i-1]
		}
	}
	for i := n-1 ; i >= 0 ; i-- {
		if i == n-1 {
			rightProduct[i] = nums[i]
		} else {
			rightProduct[i] = nums[i]*rightProduct[i+1]
		}
	}
	for i := 0 ; i < n ; i++ {
		if i == 0 {
			res = append(res, rightProduct[i+1])
		} else if i == n-1 {
			res = append(res, leftProduct[i-1])
		} else {
			res = append(res, leftProduct[i-1]*rightProduct[i+1])
		}
	}
	return res
}