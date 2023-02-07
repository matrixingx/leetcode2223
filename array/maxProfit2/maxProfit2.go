package maxprofit2

func maxProfit(prices []int) int {
	var (
		n = len(prices)
		dp = make([][]int,2)// dp0表示当天手里不含股票的最大利润，dp1表示当天手里含股票的最大利润
	)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	dp[0][0] = 0
	dp[1][0] = -prices[0]
	for i := 1 ; i < n ; i++ {
		dp[0][i] = max(dp[1][i-1]+prices[i],dp[0][i-1]) // 不含股票，可能是前几天买了今天卖，或者干脆一直不买
		dp[1][i] = max(dp[0][i-1]-prices[i],dp[1][i-1]) // 含股票，可能是前几天的股票或者当天的股票
	}
	return max(dp[0][n-1],dp[1][n-1])
}

func max(i,j int) int {
	if i > j {
		return i
	}
	return j
}