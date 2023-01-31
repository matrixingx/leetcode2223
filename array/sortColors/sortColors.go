package sortcolors

/*
	https://leetcode.cn/problems/sort-colors/
*/
// 经典的河南国旗问题
func sortColors(nums []int)  {
	var p0,p1 = 0,0 // p0代表不包括p0的p0下标前的都是0，p1代表不包括p1的p1下标前都是0和1
	for i,v := range nums {
		if v == 0 {
			nums[i],nums[p0] = nums[p0],nums[i]
			if p0 < p1 { // 代表已经有有序的1了，这时p0下标换出去的是1，需要把1换回来
				nums[i],nums[p1] = nums[p1],nums[i]
			}
			p0++
			p1++
		}
		if v == 1 {
			nums[i],nums[p1] = nums[p1],nums[i]
			p1++
		}
	}
}