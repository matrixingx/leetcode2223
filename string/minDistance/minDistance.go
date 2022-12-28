package mindistance

/*
	https://leetcode.cn/problems/edit-distance/
*/

func minDistance(word1 string, word2 string) int {
	var m,n = len(word1),len(word2)
	if m*n == 0 {
		return m+n
	}
	var dp = make([][]int,m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := m ; i >= 0 ; i-- {
		dp[i][0] = m-i
	}
	for i := 0 ; i <= n ; i++ {
		dp[m][i] = i
	}
	for i := m-1 ; i >= 0 ; i -- {
		for j := 1 ; j <= n ; j++ {
			left := dp[i][j-1]
			down := dp[i+1][j]
			leftDown := dp[i+1][j-1]
			if word1[m-1-i] == word2[j-1] {
				leftDown--
			}
			dp[i][j] = 1+min(left,down,leftDown)
		}
	}
	return dp[0][n]
}

func min(arr ...int) int {
	var min = 1<<31-1
	for i := range arr{
		if min > arr[i] {
			min = arr[i]
		}
	}
	return min
}