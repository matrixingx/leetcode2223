package uniquepaths

func uniquePaths(m int, n int) int {
	var dp = make([][]int,m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for i := range dp {
		for j := range dp[i] {
			if i == 0 {
				dp[i][j] = 1
			}
			if j == 0 {
				dp[i][j] = 1
			}
		}
	}
	for i := 1 ; i < m ; i ++ {
		for j := 1 ; j < n ; j ++ {
			dp[i][j] = dp[i-1][j]+dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}