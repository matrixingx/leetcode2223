package coinchange

func coinChange(coins []int, amount int) int {
	var dp = make([]int,amount+1)
	for i := range coins {
		if coins[i] <= amount {
			dp[coins[i]] = 1
		}
	}
	for i := 1 ; i <= amount ; i++ {
		if dp[i] == 0 {
			var minArr []int
			for _,coin := range coins {
				if i > coin && coin <= amount && dp[i-coin] > 0 {
					minArr = append(minArr, dp[i-coin]+1)
				}
			}
			if len(minArr) > 0 {
				dp[i] = min(minArr...)
			}
		}
	}
	if dp[amount] > 0 || amount == 0 {
		return dp[amount]
	}
	return -1
}

func min(arr ...int) int {
	var min = 2<<31
	for i := range arr {
		if arr[i] < min {
			min = arr[i]
		}
	}
	return min
}