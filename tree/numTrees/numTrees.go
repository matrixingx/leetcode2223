package numtrees

/*
	https://leetcode.cn/problems/unique-binary-search-trees/
*/

func numTrees(n int) int {
	var dp = make([]int,n+1)
	if n == 0 {
		return 0
	}
	dp[0],dp[1] = 1,1 // 代表0个节点，1个节点的情况
	for i := 2;i <= n ; i++ {
		for j := 1 ; j <= i ; j++ { // 代表j号做根节点的情况
			dp[i] += dp[j-1]*dp[i-j]
		}
	}
	return dp[n]
}