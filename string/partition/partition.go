package partition

/*
	https://leetcode.cn/problems/palindrome-partitioning/
*/

func partition(s string) [][]string {
	var n = len(s)
	var dp = make([][]bool,n)
	var res [][]string
	var dfs func(start int)
	var ans []string
	dfs = func(start int) {
		if start == n {
			res = append(res, append([]string{},ans...))
		}
		for i := start ; i < n ; i++ {
			if dp[start][i] {
				ans = append(ans, s[start:i+1])
				dfs(i+1)
				ans = ans[:len(ans)-1]
			}
		}
	}
	for i := range dp {
		dp[i] = make([]bool, n)
		dp[i][i] = true
	}
	for long := 1 ; long < n ; long ++ {
		var i,j = 0,0+long
		for j < n {
			if s[i] == s[j] {
				if j - i + 1 == 2 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}	
			}
			i++
			j++
		}
	}
	dfs(0)
	return res
}