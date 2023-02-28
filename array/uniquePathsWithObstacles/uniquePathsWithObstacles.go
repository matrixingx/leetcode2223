package uniquepathswithobstacles

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	var n = len(obstacleGrid)
	if n == 0 {
		return 0
	}
	var m = len(obstacleGrid[0])
	var dp = make([][]int,n)
	for i := range dp {
		dp[i] = make([]int, m)
	}
	for i := range dp {
		if obstacleGrid[i][0] == 0 {
			dp[i][0] = 1
		} else {
			break
		}
	}
	for j := range dp[0] {
		if obstacleGrid[0][j] == 0 {
			dp[0][j] = 1
		} else {
			break
		}
	}
	for i := 1 ; i < n ; i ++ {
		for j := 1 ; j < m ; j ++ {
			if obstacleGrid[i][j] == 0 {
				dp[i][j] = dp[i-1][j]+dp[i][j-1]
			} else {
				dp[i][j] = 0
			}
		}
	}
	return dp[n-1][m-1]
}