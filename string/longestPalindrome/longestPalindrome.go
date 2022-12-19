package longestpalindrome

/*
	https://leetcode.cn/problems/longest-palindromic-substring/
*/

func longestPalindrome(s string) string {
	var length = len(s)
	var dp [][]bool = make([][]bool, length)
	var res = ""
	var resLong = 1
	if len(s) == 0 {
		return res
	}
	res = s[:1]
	for i := range dp {
		dp[i] = make([]bool, length)
	}
	for i := range dp {
		dp[i][i] = true
	}
	for long := 1 ; long < len(dp) ; long++ {
		var i = 0
		var j = i + long
		for j < len(s) {
			if j-i+1 == 2 && s[i] == s[j] {
				dp[i][j] = true
			} else if j-i+1 > 2 && s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1]
			}
			if dp[i][j] && j-i+1 > resLong {
				res = s[i:j+1]
			}
			i++
			j++
		}
	}
	return res
}