package threesum

import "sort"

func threeSum(nums []int) [][]int {
	var resSet = make(map[[3]int]bool)
	sort.SliceStable(nums,func(i, j int) bool {
		return nums[i] < nums[j]
	})
	var n = len(nums)
	for i := 0 ; i < n ; i ++ {
		var j,k = i+1,n-1
		for j < k {
			if nums[i] + nums[j] + nums[k] < 0 {
				j++
			} else if nums[i] + nums[j] + nums[k] > 0 {
				k--
			} else {
				resSet[[3]int{nums[i],nums[j],nums[k]}] = true
				j++
				k--
			}
		}
	}
	var res [][]int
	for k := range resSet {
		res = append(res, []int{k[0],k[1],k[2]})
	}
	return res
}