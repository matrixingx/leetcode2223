package longestconsecutive

func longestConsecutive(nums []int) int {
	var mp = make(map[int]bool)
	for i := range nums {
		mp[nums[i]] = true
	}
	var maxCount = 0
	for i := range nums {
		if _,exist := mp[nums[i]-1];!exist {
			count := 1
			j := nums[i]+1
			_,ok2 := mp[j]
			for ok2 {
				count++
				j++
				_,ok2 = mp[j]
			}
			if count > maxCount {
				maxCount = count
			}
		}
	}
	return maxCount
}