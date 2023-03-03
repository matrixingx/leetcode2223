package test7

import "fmt"

func solution() {
	var coins = []int{1,2,5}
	var target = 13
	var dp = make([]int,target+1)
	for i := range coins {
		dp[coins[i]] = 1
	}
	for i := 1 ; i <= target ; i ++ {
		if dp[i] == 0 {
			var minValue = 1<<31
			for j := range coins {
				if i-coins[j] > 0 && dp[i-coins[j]] > 0 {
					minValue = min(minValue,dp[i-coins[j]]+1)
				}
			}
			dp[i] = minValue
		}
	}
	fmt.Println(dp[target]) 
}

func min(arr ...int) int {
	var minValue = 1<<31
	for i := range arr {
		if arr[i] < minValue {
			minValue = arr[i]
		}
	}
	return minValue
}