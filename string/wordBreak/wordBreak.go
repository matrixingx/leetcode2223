package wordbreak

/*
	https://leetcode.cn/problems/word-break/
*/

func wordBreak(s string, wordDict []string) bool {
	var mp = make(map[string]bool)
	for i := range wordDict {
		mp[wordDict[i]] = true
	}
	var n = len(s)
	var dp = make([]bool,n+1)
	dp[0] = true
	for i := 0 ; i <= n ; i ++ {
		for j := 0 ; j <= i ; j++ {
			if !dp[i] {
				dp[i] = dp[j]&&mp[string(s[j:i])]
			}
		}
	}
	return dp[n]
}