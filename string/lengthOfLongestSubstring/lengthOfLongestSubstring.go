package longestpalindrome


/*
	https://leetcode.cn/problems/longest-substring-without-repeating-characters/
*/

func lengthOfLongestSubstring(s string) int {
	if len(s) < 1 {
		return 0
	}
	var res = 1
	var rightPoint = 0
	var set = make(map[byte]bool)
	for leftPoint := 0 ; leftPoint < len(s) ; leftPoint++ {
		set[s[leftPoint]] = true
		if leftPoint == rightPoint {
			rightPoint++
		}
		for rightPoint < len(s) && !set[s[rightPoint]] {
			set[s[rightPoint]] = true
			if rightPoint-leftPoint+1 > res {
				res = rightPoint-leftPoint+1
			}
			rightPoint++
			if rightPoint == len(s) {
				return res
			}
		}
		delete(set,s[leftPoint])
	}
	return res
}