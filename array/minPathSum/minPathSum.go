package minpathsum

func minPathSum(grid [][]int) int {
	var n = len(grid)
	if n <= 0 {
		return 0
	}
	var m = len(grid[0])
	var dp = make([][]int,n)
	for i := range dp {
		dp[i] = make([]int, m)
	}
	for i := range dp {
		for j := range dp[i] {
			if i == 0 && j == 0 {
				dp[i][j] = grid[i][j]
			}
			if i == 0 && j > 0 {
				dp[i][j] += grid[i][j]+dp[i][j-1]
			}
			if j == 0 && i > 0 {
				dp[i][j] += grid[i][j]+dp[i-1][j]
			}
		}
	}
	for i := 1 ; i < n ; i ++ {
		for j := 1 ; j < m ; j ++ {
			dp[i][j] = min(dp[i-1][j],dp[i][j-1])+grid[i][j]
		}
	}
	return dp[n-1][m-1]
}

func min(i,j int) int {
	if i < j {
		return i
	}
	return j
}