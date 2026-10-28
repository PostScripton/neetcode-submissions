func maxProfit(prices []int) int {
	buy, sell, maxProfit := 0, 0, math.MinInt64

	for i := 0; i < len(prices); i++ {
		if prices[i] < prices[buy] {
			buy = i
		}
		if prices[i] > prices[sell] || sell < buy {
			sell = i
		}

		profit := prices[sell] - prices[buy]

		if profit > maxProfit {
			maxProfit = profit
		}
	}

	if maxProfit < 0 {
		maxProfit = 0
	}

	return maxProfit
}
