package strstr

func strStr(haystack string, needle string) int {
	var getNext func(needle string) []int
	getNext = func(s string) []int {
		var n = len(s)
		var next = make([]int,n+1) // next数组多开一位，0固定为-1
		next[0] = -1
		var i,j = 0,-1 // i,j为双指针下标，一开始固定next[1] = 0，即字串需要len > 1
		for i < n {
			if j == -1 || s[i] == s[j] {
				i++
				j++
				next[i] = j
			} else {
				j = next[j] // 不匹配时回溯j，最多回溯到j = next[0]即j = -1后进入if j == -1条件
			}
		}
		return next
	}
	var next = getNext(needle)
	var m,n = len(haystack),len(needle)
	var i,j = 0,0
	for i < m && j < n {
		if j == -1 || haystack[i] == needle[j] {
			i++
			j++
		} else {
			j = next[j]
		}
	}
	if j == n {
		return i-j
	}
	return -1
}