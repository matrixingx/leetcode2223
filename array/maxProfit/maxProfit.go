package maxprofit

func maxProfit(prices []int) int {
	var res = 0
	var minPrice = 1<<31
	for i := range prices {
		if prices[i] <minPrice {
			minPrice = prices[i]
		}
		if res < prices[i] - minPrice {
			res = prices[i] - minPrice
		}
	}
	return res
}
