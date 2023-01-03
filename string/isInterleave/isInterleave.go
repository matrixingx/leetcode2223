package isinterleave

/*
	https://leetcode.cn/problems/interleaving-string/
*/

func isInterleave(s1 string, s2 string, s3 string) bool {
	var m, n, l = len(s1), len(s2), len(s3)
	if m+n != l {
		return false
	}
	var dp = make([][]bool, m+1) // 0代表空串情况，所以是m+1
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true
	for i := 0 ; i <= m ; i ++ {
		for j := 0 ; j <= n ; j ++ {
			s3Idx := i + j - 1
			if i > 0 {
				dp[i][j] = dp[i][j] || (dp[i-1][j] && s1[i-1] == s3[s3Idx])
			}
			if j > 0 {
				dp[i][j] = dp[i][j] || (dp[i][j-1] && s2[j-1] == s3[s3Idx])
			}
		}
	}
	return dp[m][n]
}