package nextpermutation

func nextPermutation(nums []int) {
	var hasNext = false
	if len(nums) == 0 || len(nums) == 1 {
		return
	}
	var n = len(nums)
	var i, j = n - 2, n - 1
	for i >= 0 {
		if nums[i] >= nums[j] {
			i--
			j--
		} else {
			hasNext = true
			break
		}
	}
	if !hasNext {
		swapArr(nums,0,n-1)
		return
	}
	var k = n - 1
	for ; k > j; k-- {
		if nums[k] > nums[i] {
			break
		}
	}
	nums[i],nums[k] = nums[k],nums[i]
	swapArr(nums,j,n-1)
	return
}

func swapArr(nums []int,start,end int) {
	a,b := start,end
	for a < b   {
		nums[a],nums[b] = nums[b],nums[a]
		a++
		b--
	}
}