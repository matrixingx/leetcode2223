package numsquares

func numSquares(n int) int {
	var dp = make([]int,n+1)
	var minValue = 1<<31
	for i := 1 ; i <= n ; i ++ {
		m := minValue
		for j := 1 ; j*j<=i ; j++ {
			m = min(m,dp[i-j*j])
		}
		dp[i] = m+1
	}
	return dp[n]
}

func min(i,j int) int {
	if i < j {
		return i
	}
	return j
}