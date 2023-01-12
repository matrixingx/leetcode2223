package threesumclosest

import "sort"

func threeSumClosest(nums []int, target int) int {
	var ans = 1<<31 - 1
	var n = len(nums)
	var res []int
	sort.SliceStable(nums,func(i, j int) bool {
		return nums[i] < nums[j]
	})
	for i := 0 ; i < n ; i ++ {
		var j,k = i+1,n-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			absoluteRes := absolute(sum,target)
			if absoluteRes < ans {
				ans = absoluteRes
				res = []int{nums[i] , nums[j] , nums[k]}
			}
			if sum < target {
				j++
			} else if sum > target {
				k--
			} else {
				break
			}
		}
	}
	return res[0]+res[1]+res[2]
}

func absolute(a,b int) int {
	if a - b >= 0 {
		return a-b
	}
	return b-a
}