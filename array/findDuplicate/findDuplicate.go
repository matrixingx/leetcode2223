package findduplicate

func findDuplicate(nums []int) int {
	if len(nums) < 2 {
		return -1
	}
	slow,fast := nums[0],nums[nums[0]]
	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}
	q := fast
	p := 0
	for p != q {
		p = nums[p]
		q = nums[q]
	}
	return p
}