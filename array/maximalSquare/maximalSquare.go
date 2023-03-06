package maximalsquare

func maximalSquare(matrix [][]byte) int {
	var n = len(matrix)
	var res = 0
	if n == 0 {
		return 0
	}
	var m = len(matrix[0])
	var dp = make([][]int,n)
	for i := range dp {
		dp[i] = make([]int, m)
	}
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == '1' {
				dp[i][j] = 1
				res = 1
			}
		}
	}
	for i := 1 ; i < n ; i++ {
		for j := 1 ; j < m ; j ++ {
			if dp[i][j] > 0 {
				dp[i][j] = min(dp[i-1][j-1]+1,dp[i-1][j]+1,dp[i][j-1]+1)
			}
			if dp[i][j] > res {
				res = dp[i][j]
			}
		}
	}
	return res*res
}

func min(arr ...int) int {
	if len(arr) == 0 {
		return 0
	}
	var min = 1<<31
	for i := range arr {
		if arr[i] < min {
			min = arr[i]
		}
	}
	return min
}